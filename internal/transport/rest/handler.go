package rest

import (
	"encoding/json"
	"github.com/TarasTarkovskyi/CRUD-6-template/internal/domain"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"path"
	"strconv"
)

type Persons interface {
	Create(person domain.PersonInfo)
	GetByID(id int) domain.PersonInfo
	GetAll() []domain.PersonInfo
	Update(id int, input domain.UpdatePersonInput)
	Delete(id int)
}

type Handler struct {
	personsService Persons
}

func NewHandler(p Persons) Handler {
	return Handler{personsService: p}
}

func (h Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	var person domain.PersonInfo
	json.Unmarshal(body, &person)

	h.personsService.Create(person)
}

func (h Handler) GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	person := h.personsService.GetByID(id)

	jsonMessage, _ := json.Marshal(person)
	w.Write(jsonMessage)
}

func (h Handler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	books := h.personsService.GetAll()
	jsonMessage, _ := json.Marshal(books)
	w.Write(jsonMessage)
}

func (h Handler) UpdateByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))
	requestBody, _ := io.ReadAll(r.Body)

	var input domain.UpdatePersonInput
	json.Unmarshal(requestBody, &input)

	h.personsService.Update(id, input)
}

func (h Handler) DeleteByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	h.personsService.Delete(id)
}

func (h Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/person").Subrouter()

	s.HandleFunc("", h.CreateHandler).Methods(http.MethodPost)
	s.HandleFunc("/{id}", h.GetByIDHandler).Methods(http.MethodGet)
	s.HandleFunc("", h.GetAllHandler).Methods(http.MethodGet)
	s.HandleFunc("/{id}", h.UpdateByIDHandler).Methods(http.MethodPut)
	s.HandleFunc("/{id}", h.DeleteByIDHandler).Methods(http.MethodDelete)

	return s
}
