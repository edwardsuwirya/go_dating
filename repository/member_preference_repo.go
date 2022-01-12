package repository

import (
	"context"
	"github.com/edwardsuwirya/go_dating/entity"
	"github.com/edwardsuwirya/go_dating/util/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type MemberPreferenceRepo interface {
	Create(newMemberPref *entity.MemberPreferences, newMemberInterest []entity.MemberInterest) error
}

type memberPreferenceRepo struct {
	conn *pgx.Conn
}

func (m *memberPreferenceRepo) Create(newMemberPref *entity.MemberPreferences, newMemberInterest []entity.MemberInterest) error {
	logger.Log.Debug().Msgf("Create %v, interest %v", newMemberPref, newMemberInterest)
	tx, err := m.conn.BeginTx(context.Background(), pgx.TxOptions{})
	defer func(tx pgx.Tx) {
		if err != nil {
			logger.Log.Error().Err(err).Msg("Failed create preference member")
			tx.Rollback(context.Background())
		} else {
			tx.Commit(context.Background())
		}
	}(tx)

	_, err = tx.Exec(context.Background(), "insert into member_preference values($1,$2,$3,$4,$5,$6)",
		uuid.New().String(),
		newMemberPref.MemberId,
		newMemberPref.LookingForGender,
		newMemberPref.LookingForDomicile,
		newMemberPref.LookingForStartAge,
		newMemberPref.LookingForEndAge)

	for _, i := range newMemberInterest {
		_, err = tx.Exec(context.Background(), "insert into member_interest values($1,$2)",
			i.InterestId,
			i.MemberId)
	}
	if err != nil {
		return err
	}
	return nil
}

func NewMemberPreferenceRepo(conn *pgx.Conn) MemberPreferenceRepo {
	return &memberPreferenceRepo{
		conn: conn,
	}
}
