package reservation

import "github.com/canyouhearthemusic/kamigr/db/sqlc"

type Service struct {
	dao *sqlc.Queries
}

func New(dao *sqlc.Queries) *Service {
	return &Service{
		dao: dao,
	}
}
