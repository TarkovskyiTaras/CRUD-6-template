package main

import (
	"github.com/TarasTarkovskyi/CRUD-6-template/internal/repository/db_psql"
	"github.com/TarasTarkovskyi/CRUD-6-template/internal/service"
	"github.com/TarasTarkovskyi/CRUD-6-template/internal/transport/rest"
	"github.com/TarasTarkovskyi/CRUD-6-template/pkg/database"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	//dbMap := make(map[int]domain.PersonInfo)
	//personsRepo := db_map.NewPersons(dbMap)

	dbPsql := database.NewPostgresConnection(database.ConnectionInfo{Host: "localhost", Port: 5432, UserName: "crud-6", DBName: "crud-6-db", SSLMode: "disable", Password: "12345"})
	personsRepo := db_psql.NewPersons(dbPsql)
	personsService := service.NewPersons(personsRepo)
	handler := rest.NewHandler(personsService)
	//handler := printout.NewHandler(personsService)

	//defer dbPsql.Close()

	server := http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	server.ListenAndServe()

	//handler.CreateHandler(domain.PersonInfo{ID: 1, FirstName: "taras", LastName: "t", DOB: "1992-01-23T00:00:00Z", AddressAndPhone: "093"})
	//handler.GetByIDHandler(3)
}
