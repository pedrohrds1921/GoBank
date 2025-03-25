package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/pedrohrds1921/GoBank/internal/services"
)

type Api struct {
	Router         *chi.Mux
	PersonServices services.PersonServices
	CompanyService services.CompanyServices
}

func (a *Api) Routes() {

	a.Router.Route("/person", func(r chi.Router) {
		r.Post("/conta", a.HandleCreateAccount)
		r.Get("/conta/{id}/saldo", a.HandleCheckBalance)
		r.Post("/conta/{id}/saque", a.HandleWithDrawAccount)
		r.Post("/conta/{id}/deposito", a.HandleDepositBalance)
	})
}
