package repository

import (
	"database/sql"

	"github.com/Gena97/telegram_bot/internal/app/model"
	"github.com/Gena97/telegram_bot/internal/service"
)

type PGXRepository struct {
	mainDB *sql.DB
}

type PGXRepo interface {
	Hashes
}

type Users interface {
	GetUserInfosMap() (map[int64]model.User, error)
}

type Hashes interface {
	GetFileHashes(lenHashes int, file_type string) ([]service.FileHash, error)
	InsertHash(fileHash service.FileHash) error
}

func InitDatabase() (PGXRepository, error) {
	var err error
	var repo PGXRepository
	repo.mainDB, err = sql.Open("sqlite", "../../database/sql_lite.db")
	if err != nil {
		return PGXRepository{}, err
	}
	return repo, nil
}
