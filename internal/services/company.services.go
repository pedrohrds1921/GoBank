package services

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pedrohrds1921/GoBank/internal/repository"
)

type CompanyServices struct {
	SqlConn    *pgxpool.Pool
	repository repository.Queries
}

func NewCompanyServices(SqlConn *pgxpool.Pool) CompanyServices {
	return CompanyServices{
		SqlConn:    SqlConn,
		repository: *repository.New(SqlConn),
	}
}
