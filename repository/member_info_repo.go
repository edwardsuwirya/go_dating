package repository

import (
	"context"
	"github.com/edwardsuwirya/go_dating/entity"
	"github.com/edwardsuwirya/go_dating/util/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type MemberInfoRepo interface {
	FindById(memberId string) (*entity.Member, error)
	Create(newMember *entity.Member) (*entity.Member, error)
}

type memberInfoRepo struct {
	conn *pgx.Conn
}

func (m *memberInfoRepo) FindById(memberId string) (*entity.Member, error) {
	logger.Log.Debug().Msgf("Find by %s", memberId)
	var member entity.Member
	err := m.conn.QueryRow(context.Background(), `
	SELECT pi.member_id,pi.self_description,pi.name,pi.photo_path, ci.instagram_id,ci.twitter_id,ai.city 
	FROM
		member_personal_information pi JOIN member_address_information ai ON pi.member_id = ai.member_id
		JOIN member_contact_information ci ON pi.member_id = ci.member_id
	WHERE pi.member_id = $1;
	`, memberId).Scan(
		&member.PersonalInfo.MemberId,
		&member.PersonalInfo.SelfDescription,
		&member.PersonalInfo.Name,
		&member.PersonalInfo.RecentPhotoPath,
		&member.ContactInfo.InstagramId,
		&member.ContactInfo.TwitterId,
		&member.AddressInfo.City)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (m *memberInfoRepo) Create(newMember *entity.Member) (*entity.Member, error) {
	logger.Log.Debug().Msgf("Create %s", newMember)
	tx, err := m.conn.BeginTx(context.Background(), pgx.TxOptions{})
	defer func(tx pgx.Tx) {
		if err != nil {
			logger.Log.Error().Err(err).Msg("Failed update profile member")
			tx.Rollback(context.Background())
		} else {
			tx.Commit(context.Background())
		}
	}(tx)
	_, err = tx.Exec(context.Background(),
		"insert into member_personal_information(personal_information_id,member_id,name,bod,gender,photo_path,self_description) values($1,$2,$3,$4,$5,$6,$7)",
		uuid.New().String(), newMember.PersonalInfo.MemberId, newMember.PersonalInfo.Name, newMember.PersonalInfo.Bod, newMember.PersonalInfo.Gender, "", newMember.PersonalInfo.SelfDescription)
	_, err = tx.Exec(context.Background(), "insert into member_address_information(address_information_id,member_id,address,city,postal_code) values($1,$2,$3,$4,$5)",
		uuid.New().String(), newMember.PersonalInfo.MemberId, newMember.AddressInfo.Address, newMember.AddressInfo.City, newMember.AddressInfo.PostalCode)
	_, err = tx.Exec(context.Background(), "update member_contact_information set mobile_phone = $2, instagram_id = $3, twitter_id = $4 where member_id=$1",
		newMember.PersonalInfo.MemberId, newMember.ContactInfo.MobilePhoneNumber, newMember.ContactInfo.InstagramId, newMember.ContactInfo.TwitterId)

	if err != nil {
		return nil, err
	}
	return newMember, nil
}

func NewMemberPersonalInfoRepo(conn *pgx.Conn) MemberInfoRepo {
	memberPersonalInfoRepo := &memberInfoRepo{
		conn: conn,
	}
	return memberPersonalInfoRepo
}
