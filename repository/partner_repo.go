package repository

import (
	"context"
	"errors"
	"github.com/edwardsuwirya/go_dating/entity"
	"github.com/edwardsuwirya/go_dating/util/logger"
	"github.com/jackc/pgx/v4"
)

type PartnerRepo interface {
	Find(id string, byGender string, byDomicile string, byStartAge int, byEndPage int, byInterest []string, limit int, offset int) ([]entity.MemberPersonalInformation, error)
	Create(id string, id2 string) error
	FindAll(id string) ([]entity.MemberPersonalInformation, error)
}

type partnerRepo struct {
	conn *pgx.Conn
}

func (p *partnerRepo) Find(id string, byGender string, byDomicile string, byStartAge int, byEndPage int, byInterest []string, limit int, pageNo int) ([]entity.MemberPersonalInformation, error) {
	logger.Log.Debug().Msgf("Find id:%s, gender:%s, domicile:%s, startAge:%d, endAge:%d", id, byGender, byDomicile, byStartAge, byEndPage)
	rows, err := p.conn.Query(context.Background(), `
		SELECT pi.member_id,pi.name,pi.photo_path,pi.self_description
		from member_personal_information pi join member_preference p
 		on pi.member_id = p.member_id AND
		date_part('year', age(now(), pi.bod)) BETWEEN $4 AND $5 AND 
 		p.member_id != $1 AND
		p.looking_gender <> $2 AND 			
		p.looking_domicile = $3 
		join member_interest i 
 			on p.member_id = i.member_id AND 
 			i.interest_id =ANY($8)
		order by pi.member_id        
		limit $6 offset $7
`, id, byGender, byDomicile, byStartAge, byEndPage, limit, (pageNo*limit)-limit, byInterest)
	if err != nil {
		return nil, err
	}
	var members []entity.MemberPersonalInformation
	for rows.Next() {
		var member entity.MemberPersonalInformation
		rows.Scan(&member.MemberId, &member.Name, &member.RecentPhotoPath, &member.SelfDescription)
		members = append(members, member)
	}
	logger.Log.Debug().Msgf("Result %v", members)
	return members, nil
}

func (p *partnerRepo) Create(id string, id2 string) error {
	logger.Log.Debug().Msgf("Create new partner for %s", id)
	res, err := p.conn.Exec(context.Background(), "insert into member_partner values($1,$2)", id, id2)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed insert member_partner")
		return err
	}
	if res.RowsAffected() == 0 {
		logger.Log.Error().Msg("Failed insert member_partner")
		return errors.New("Insert Failed")
	}
	return nil
}

func (p *partnerRepo) FindAll(id string) ([]entity.MemberPersonalInformation, error) {
	logger.Log.Debug().Msgf("Find id:%s", id)
	rows, err := p.conn.Query(context.Background(), `
		select pi.member_id,pi.name,pi.photo_path,pi.self_description from 
		member_personal_information pi join (
		select partner_id from member_partner where member_id = $1
		) p on pi.member_id = p.partner_id;
`, id)
	if err != nil {
		return nil, err
	}
	var members []entity.MemberPersonalInformation
	for rows.Next() {
		var member entity.MemberPersonalInformation
		rows.Scan(&member.MemberId, &member.Name, &member.RecentPhotoPath, &member.SelfDescription)
		members = append(members, member)
	}
	logger.Log.Debug().Msgf("Result %v", members)
	return members, nil
}

func NewPartnerRepo(conn *pgx.Conn) PartnerRepo {
	return &partnerRepo{
		conn: conn,
	}
}
