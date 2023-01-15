package db_map

import (
	"CRUD-6-template/internal/domain"
	"fmt"
)

type Persons struct {
	db map[int]domain.PersonInfo
}

func NewPersons(db map[int]domain.PersonInfo) Persons {
	return Persons{db}
}

func (p Persons) Create(person domain.PersonInfo) {
	p.db[person.ID] = person
}

func (p Persons) GetByID(id int) domain.PersonInfo {
	return p.db[id]
}
func (p Persons) GetAll() []domain.PersonInfo {
	all := make([]domain.PersonInfo, 10)
	for i, c := range p.db {
		all[i] = c
	}
	return all
}

func (p Persons) Update(id int, person domain.UpdatePersonInput) {
	if _, ok := p.db[id]; !ok {
		fmt.Println("no such id in the database yet")
		return
	}
	updatedPerson := domain.PersonInfo{
		ID: id,
	}

	if person.FirstName != nil {
		updatedPerson.FirstName = *person.FirstName
	}

	if person.LastName != nil {
		updatedPerson.LastName = *person.LastName
	}
	if person.DOB != nil {
		updatedPerson.DOB = *person.DOB
	}
	if person.HomeAddress != nil {
		updatedPerson.HomeAddress = *person.HomeAddress
	}
	if person.CellPhone != nil {
		updatedPerson.CellPhone = *person.CellPhone
	}

	p.db[id] = updatedPerson

	fmt.Println(p.db)
}

func (p Persons) Delete(id int) {
	delete(p.db, id)
}
