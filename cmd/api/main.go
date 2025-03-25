package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/pedrohrds1921/GoBank/internal/api"
	"github.com/pedrohrds1921/GoBank/internal/services"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	ctx := context.Background()
	SqlConn, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("GOBANK_DATABASE_USER"),
		os.Getenv("GOBANK_DATABASE_PASSWORD"),
		os.Getenv("GOBANK_DATABASE_HOST"),
		os.Getenv("GOBANK_DATABASE_PORT"),
		os.Getenv("GOBANK_DATABASE_NAME"),
	))
	if err != nil {
		panic(err)
	}
	defer SqlConn.Close()
	api := api.Api{
		Router:         chi.NewMux(),
		PersonServices: services.NewPersonServices(SqlConn),
		CompanyService: services.NewCompanyServices(SqlConn),
	}
	api.Routes()
	fmt.Println("Start Server on port :3000")
	if err := http.ListenAndServe("localhost:3000", api.Router); err != nil {
		panic(err)
	}
}
