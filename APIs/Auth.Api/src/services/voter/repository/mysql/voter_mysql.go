package mysql

import (
	"context"
	"database/sql"
	"log"

	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/config"
	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/models"
)

type voterMysqlRepository struct {
	DB *sql.DB
}

func NewVoterMysqlRepository(db *sql.DB) models.IVoterMysqlRepository {
	return &voterMysqlRepository{
		DB: db,
	}
}

// CreateVoter implements models.IVoterMysqlRepository
func (r *voterMysqlRepository) CreateVoter(ctx context.Context, nik int64, nonce int64, previousHash string, hash string) error {
	tx, err := r.DB.Begin()
	if err != nil {
		log.Printf("Error Begin Transaction InsertUser : %v\n", err)
		return err
	}
	defer config.CommitOrRollback(tx)

	query := `INSERT INTO voters (userId, nonce, prevHash, hash) VALUES (?,?,?,?)`
	_, err = r.DB.ExecContext(ctx, query, nik, nonce, previousHash, hash)
	if err != nil {
		log.Printf("Error when inserting data to db : %v\n", err)
		return err
	}
	return nil
}

// GetVotersByNIK implements models.IVoterMysqlRepository
func (r *voterMysqlRepository) GetVotersByNIK(ctx context.Context, nik int64) (*models.Voter, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		log.Printf("Error Begin Transaction GetVotersByNIK : %v\n", err)
		return nil, err
	}
	defer config.CommitOrRollback(tx)

	query := "SELECT userId, nonce, prevHash, hash FROM voters WHERE userId=?"
	rows, err := r.DB.Query(query, nik)
	if err != nil {
		log.Printf("Error GetVotersByNIK : %v\n", err)
		return nil, err

	}
	res := []models.Voter{}

	for rows.Next() {
		user := models.Voter{}
		err = rows.Scan(&user.UserId, &user.Nonce, &user.PreviousHash, &user.Hash)
		if err != nil {
			log.Printf("Error when Scanning data GetVotersByNIK : %v", err)
			return nil, err
		}
		res = append(res, user)
	}
	if len(res) == 0 {
		return nil, nil
	}
	return &res[0], nil
}
