package printout

import (
	"fmt"
	"github.com/TarasTarkovskyi/CRUD-6-template/internal/domain"
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

func NewHandler(ps Persons) Handler {
	return Handler{personsService: ps}
}

func (h Handler) CreateHandler(p domain.PersonInfo) {
	h.personsService.Create(p)
	fmt.Println("User successfully created")
}

func (h Handler) GetByIDHandler(id int) {
	person := h.personsService.GetByID(id)
	fmt.Println(person)
}

func (h Handler) UpdateByIDHandler(id int, p domain.UpdatePersonInput) {
	h.personsService.Update(id, p)
	fmt.Println("User successfully updated")
}

func (h Handler) DeleteByIDHandler(id int) {
	h.personsService.Delete(id)
	fmt.Println("User successfully deleted")
}
