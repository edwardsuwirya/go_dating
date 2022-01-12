package repository

import (
	"context"
	"errors"
	"github.com/edwardsuwirya/go_dating/entity"
	"github.com/edwardsuwirya/go_dating/util/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type MemberAccessRepo interface {
	FindByUserNameAndPasswordAndVerified(userName string, password string) (bool, error)
	UpdateVerification(id string) error
	Create(newAccess *entity.MemberUserAccess) error
}

type memberAccessRepo struct {
	conn *pgx.Conn
}

func (m *memberAccessRepo) FindByUserNameAndPasswordAndVerified(userName string, password string) (bool, error) {
	logger.Log.Debug().Msgf("Authenticate %s", userName)
	var memberAccess entity.MemberUserAccess
	err := m.conn.QueryRow(context.Background(), "select member_id from member_access where verification_status='Y' and user_name=$1 and user_password=$2", userName, password).Scan(&memberAccess.MemberId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *memberAccessRepo) UpdateVerification(id string) error {
	logger.Log.Debug().Msgf("Update Verification %s", id)
	tx, err := m.conn.BeginTx(context.Background(), pgx.TxOptions{})
	defer func(tx pgx.Tx) {
		if err != nil {
			logger.Log.Error().Err(err).Msg("Failed update verification member")
			tx.Rollback(context.Background())
		} else {
			tx.Commit(context.Background())
		}
	}(tx)
	res, err := tx.Exec(context.Background(), "update member_access set verification_status = 'Y' where member_id=$1", id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("Update verification failed")
	}
	var email string
	err = tx.QueryRow(context.Background(), "select user_name from member_access where member_id=$1", id).Scan(&email)
	if err != nil {
		return err
	}
	_, err = tx.Exec(context.Background(), "insert into member_contact_information(contact_information_id,member_id,email) values($1,$2,$3)", uuid.New().String(), id, email)
	return nil
}

func (m *memberAccessRepo) Create(newAccess *entity.MemberUserAccess) error {
	logger.Log.Debug().Msgf("Create %s", newAccess.UserName)
	res, err := m.conn.Exec(context.Background(), "insert into member_access values($1,$2,$3,$4,$5)", newAccess.MemberId, newAccess.UserName, newAccess.Password, newAccess.JoinDate, newAccess.VerificationStatus)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed insert member_access")
		return err
	}
	if res.RowsAffected() == 0 {
		logger.Log.Error().Msg("Failed insert member_access")
		return errors.New("Insert Failed")
	}
	return nil
}

func NewMemberAccessRepo(conn *pgx.Conn) MemberAccessRepo {
	memberAccessRepo := &memberAccessRepo{
		conn: conn,
	}
	return memberAccessRepo
}
