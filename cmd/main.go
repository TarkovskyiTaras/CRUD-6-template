package main

import (
	"CRUD-6-template/internal/domain"
	"CRUD-6-template/internal/repository/db_map"
	"CRUD-6-template/internal/service"
	"CRUD-6-template/internal/transport/rest"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	dbMap := make(map[int]domain.PersonInfo)
	personsRepo := db_map.NewPersons(dbMap)

	//dbPsql := database.NewPostgresConnection(database.ConnectionInfo{
	//	Host:     "localhost",
	//	Port:     5432,
	//	UserName: "postgres",
	//	DBMame:   "persons",
	//	SSLMode:  "disable",
	//	Password: "1234",
	//})
	//defer dbPsql.Close()

	//personsRepo := db_psql.NewPersons(dbPsql)
	personsService := service.NewPersons(personsRepo)
	handler := rest.NewHandler(personsService)
	//handler := printout.NewHandler(personsService)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler.InitRouter(),
	}

	server.ListenAndServe()

	//personsRepo := db_map.Persons{ // bookrepo
	//	db: dbMap,
	//}
	//personsService := service.Persons{ // bookservice
	//	repo: personsRepo,
	//}
	//handler := rest.Handler{ // handler
	//	personsService: personsService,
	//}

	//handler := printout.Handler{
	//	db: personsRepo,
	//}
	//handler.CreateHandler(domain.PersonInfo{ID: 1, FirstName: "taras", LastName: "t", DOB: "jan 23", AddressAndPhone: "093"})
	//handler.GetByIDHandler(3)
}
