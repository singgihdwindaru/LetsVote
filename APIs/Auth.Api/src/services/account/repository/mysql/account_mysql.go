package mysql

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/config"
	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/models"
)

type accountMysqlRepository struct {
	DB *sql.DB
}

func NewAccountMysqlRepository(db *sql.DB) models.IUserMysqlRepository {
	return &accountMysqlRepository{
		DB: db,
	}
}

// GetUserByNIK implements models.IUserMysqlRepository
func (r *accountMysqlRepository) GetUserByNIK(ctx context.Context, nik int64) (*models.User, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		log.Printf("Error Begin Transaction GetUserByNIK : %v\n", err)
		return nil, err
	}
	defer config.CommitOrRollback(tx)

	query := "SELECT id, guid, metadata, hash FROM users WHERE id=?"
	rows, err := r.DB.Query(query, nik)
	if err != nil {
		log.Printf("Error GetUserByNIK : %v\n", err)
		return nil, err

	}
	res := []models.User{}

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.Id, &user.Guid, &user.Metadata, &user.Hash)
		if err != nil {
			log.Printf("Error when Scanning data GetUserByNIK : %v", err)
			return nil, err
		}
		res = append(res, user)
	}
	if len(res) == 0 {
		return nil, nil
	}
	return &res[0], nil
}

// InsertUser implements models.IUserMysqlRepository
func (r *accountMysqlRepository) InsertUser(ctx context.Context, guid, metadata, hash string) error {
	tx, err := r.DB.Begin()
	if err != nil {
		log.Printf("Error Begin Transaction InsertUser : %v\n", err)
		return err
	}
	defer config.CommitOrRollback(tx)
	var bindMeta map[string]interface{}
	_ = json.Unmarshal([]byte(metadata), &bindMeta) // TODO
	nik, ok := bindMeta["nik"].(float64)
	if !ok {
		log.Printf("Error when inserting data to db : %v\n", err)
		return err
	}

	query := `INSERT INTO users (id, guid, metadata, hash) VALUES (?,?,?,?)`
	_, err = r.DB.ExecContext(ctx, query, nik, guid, metadata, hash)
	if err != nil {
		log.Printf("Error when inserting data to db : %v\n", err)
		return err
	}
	return nil
}
