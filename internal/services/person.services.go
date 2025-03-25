package services

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pedrohrds1921/GoBank/internal/repository"
)

type PersonServices struct {
	SqlConn    *pgxpool.Pool
	repository repository.Queries
}

type CreatePersonAccountParams struct {
	MonthlyIncome float64 `json:"monthly_income"`
	Age           int32   `json:"age"`
	FullName      string  `json:"full_name"`
	Phone         string  `json:"phone"`
	Email         string  `json:"email"`
	Category      string  `json:"category"`
	Balance       float64 `json:"balance"`
}

type UpdatePersonBalanceRequest struct {
	WithdrawValue float64 `json:"WithdrawValue"`
}
type UpdatePersonDepositRequest struct {
	DepositValue float64 `json:"DepositValue"`
}

func NewPersonServices(SqlConn *pgxpool.Pool) PersonServices {
	return PersonServices{
		SqlConn:    SqlConn,
		repository: *repository.New(SqlConn),
	}
}

func (ps *PersonServices) CreatePersonAccount(ctx context.Context, data repository.CreatePersonAccountParams) (repository.Person, error) {

	account, err := ps.repository.CreatePersonAccount(ctx, data)
	if err != nil {
		return repository.Person{}, err
	}

	return account, nil
}

var ErrInsufficientFunds = errors.New("insufficient balance")

func (ps *PersonServices) CheckBalance(ctx context.Context, id uuid.UUID) (repository.Person, error) {
	account, err := ps.repository.GetPersonAcountById(ctx, id)
	if err != nil {
		return repository.Person{}, err
	}
	return account, err

}

func (ps *PersonServices) DepositValue(ctx context.Context, id uuid.UUID, depositValoue float64) error {
	account, err := ps.repository.GetPersonAcountById(ctx, id)
	if err != nil {
		return err
	}
	accountBalance, _ := account.Balance.Float64Value()
	newBalance := accountBalance.Float64 + depositValoue
	var newBalanceNumeric pgtype.Numeric
	if err := newBalanceNumeric.Scan(strconv.FormatFloat(newBalance, 'f', -1, 64)); err != nil {
		return fmt.Errorf("erro ao definir novo saldo: %w", err)
	}
	if err := ps.repository.UpdatePersonBalance(ctx, repository.UpdatePersonBalanceParams{
		ID:      id,
		Balance: newBalanceNumeric,
	}); err != nil {
		return err
	}
	return nil
}
func (ps *PersonServices) WithdrawBalance(ctx context.Context, id uuid.UUID, withdrawValue float64) error {
	account, err := ps.repository.GetPersonAcountById(ctx, id)
	if err != nil {
		return err
	}
	accountBalance, _ := account.Balance.Float64Value()
	if accountBalance.Float64 < withdrawValue {
		log.Println("aqui")
		return ErrInsufficientFunds
	}
	newBalance := accountBalance.Float64 - withdrawValue
	var newBalanceNumeric pgtype.Numeric
	if err := newBalanceNumeric.Scan(strconv.FormatFloat(newBalance, 'f', -1, 64)); err != nil {
		return fmt.Errorf("erro ao definir novo saldo: %w", err)
	}
	if err := ps.repository.UpdatePersonBalance(ctx, repository.UpdatePersonBalanceParams{
		ID:      id,
		Balance: newBalanceNumeric,
	}); err != nil {
		return err
	}
	return nil
}
