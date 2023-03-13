package service

import (
	"github.com/TarasTarkovskyi/CRUD-6-template/internal/domain"
)

type PersonsRepository interface {
	Create(person domain.PersonInfo)
	GetByID(id int) domain.PersonInfo
	GetAll() []domain.PersonInfo
	Update(id int, input domain.UpdatePersonInput)
	Delete(id int)
}

type Persons struct {
	repo PersonsRepository
}

func NewPersons(p PersonsRepository) Persons {
	return Persons{p}
}

func (p Persons) Create(person domain.PersonInfo) {
	p.repo.Create(person)
}

func (p Persons) GetByID(id int) domain.PersonInfo {
	return p.repo.GetByID(id)
}

func (p Persons) GetAll() []domain.PersonInfo {
	return p.repo.GetAll()
}

func (p Persons) Update(id int, input domain.UpdatePersonInput) {
	p.repo.Update(id, input)
}

func (p Persons) Delete(id int) {
	p.repo.Delete(id)
}

//type Persons struct {
//	repo PersonsRepository
//}
//
//func Create(p Persons, person domain.PersonInfo) {
//	p.repo.Create(person)
//}
//
//func GetByID(p Persons, id int) domain.PersonInfo {
//	return p.repo.GetByID(id)
//}
//
//func Update(p Persons, id int, person domain.PersonInfo) {
//	p.repo.Update(id, person)
//}
//
//func Delete(p Persons, id int) {
//	p.repo.Delete(id)
//}
