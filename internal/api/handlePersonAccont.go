package api

import (
	"errors"
	"net/http"

	z "github.com/Oudwins/zog"
	zhttp "github.com/Oudwins/zog/zhttp"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	jsonutils "github.com/pedrohrds1921/GoBank/internal/JsonUtils"
	"github.com/pedrohrds1921/GoBank/internal/repository"
	"github.com/pedrohrds1921/GoBank/internal/services"
	"github.com/pedrohrds1921/GoBank/internal/utils"
)

func (a *Api) HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var body repository.CreatePersonAccountParams
	if err := utils.PersonAccountRequestValidador.Parse(zhttp.Request(r), &body); err != nil {
		jsonutils.SendJson(w, z.Issues.SanitizeMap(err), http.StatusUnprocessableEntity)
		return
	}
	_, err := a.PersonServices.CreatePersonAccount(r.Context(), body)
	if err != nil {
		jsonutils.SendJson(w, err, http.StatusUnprocessableEntity)
		return
	}
	jsonutils.SendJson(w, "Conta criada com successo!", http.StatusCreated)
}

func (a *Api) HandleCheckBalance(w http.ResponseWriter, r *http.Request) {

	idRequest := chi.URLParam(r, "id")

	account, err := a.PersonServices.CheckBalance(r.Context(), uuid.MustParse(idRequest))

	if err != nil {
		jsonutils.SendJson(w, "Conta nao encontrada", http.StatusNoContent)
		return
	}

	jsonutils.SendJson(w, map[string]any{
		"Saldo é ": account.Balance,
	}, http.StatusOK)
}

func (a *Api) HandleWithDrawAccount(w http.ResponseWriter, r *http.Request) {
	idRequest := chi.URLParam(r, "id")
	var body services.UpdatePersonBalanceRequest
	if err := utils.PersonWithdrawRequestValidador.Parse(zhttp.Request(r), &body); err != nil {
		jsonutils.SendJson(w, z.Issues.SanitizeMap(err), http.StatusUnprocessableEntity)
		return
	}
	err := a.PersonServices.WithdrawBalance(r.Context(), uuid.MustParse(idRequest), body.WithdrawValue)
	if err != nil {
		if errors.Is(err, services.ErrInsufficientFunds) {
			jsonutils.SendJson(w, "Saldo insuficiente", http.StatusOK)
			return
		}
		jsonutils.SendJson(w, "Internal server erro", http.StatusInternalServerError)
		return
	}
	jsonutils.SendJson(w, "Saque Realizado", http.StatusOK)
}

func (a *Api) HandleDepositBalance(w http.ResponseWriter, r *http.Request) {
	idRequest := chi.URLParam(r, "id")
	var body services.UpdatePersonDepositRequest
	if err := utils.PersonDepositRequest.Parse(zhttp.Request(r), &body); err != nil {
		jsonutils.SendJson(w, z.Issues.SanitizeMap(err), http.StatusUnprocessableEntity)
		return
	}
	err := a.PersonServices.DepositValue(r.Context(), uuid.MustParse(idRequest), body.DepositValue)
	if err != nil {
		jsonutils.SendJson(w, "Internal server erro", http.StatusInternalServerError)
		return
	}
	jsonutils.SendJson(w, "Deposito realizado ", http.StatusOK)
}
