package db_psql

import (
	"CRUD-6-template/internal/domain"
	"database/sql"
	"fmt"
	"strings"
)

type Persons struct {
	db *sql.DB
}

func NewPersons(db *sql.DB) Persons {
	return Persons{db: db}
}

func (p Persons) Create(person domain.PersonInfo) {
	p.db.Exec("INSERT INTO persons (id, first_name, last_name, dob, home_address, cellphone) VALUES ($1, $2, $3, $4, $5, $6)", person.ID,
		person.FirstName, person.LastName, person.DOB, person.HomeAddress, person.CellPhone)
}

func (p Persons) GetByID(id int) domain.PersonInfo {
	var person domain.PersonInfo
	p.db.QueryRow("SELECT id, first_name, last_name, dob, home_address, cellphone FROM persons WHERE id = $1", id).
		Scan(&person.ID, &person.FirstName, &person.LastName, &person.DOB, &person.HomeAddress, &person.CellPhone)
	return person
}

func (p Persons) GetAll() []domain.PersonInfo {
	rows, _ := p.db.Query("SELECT id, first_name, last_name, dob, home_address, cellphone FROM persons")
	persons := make([]domain.PersonInfo, 0)
	for rows.Next() {
		var person domain.PersonInfo
		rows.Scan(&person.ID, &person.FirstName, &person.LastName, &person.DOB, &person.HomeAddress, &person.CellPhone)
		persons = append(persons, person)
	}
	return persons
}

func (p Persons) Update(id int, input domain.UpdatePersonInput) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.FirstName != nil {
		setValues = append(setValues, fmt.Sprintf("first_name=$%d", argID))
		args = append(args, *input.FirstName)
		argID++
	}

	if input.LastName != nil {
		setValues = append(setValues, fmt.Sprintf("last_name=$%d", argID))
		args = append(args, *input.LastName)
		argID++
	}

	if input.DOB != nil {
		setValues = append(setValues, fmt.Sprintf("dob=$%d", argID))
		args = append(args, *input.DOB)
		argID++
	}

	if input.HomeAddress != nil {
		setValues = append(setValues, fmt.Sprintf("home_address=$%d", argID))
		args = append(args, *input.HomeAddress)
		argID++
	}

	if input.CellPhone != nil {
		setValues = append(setValues, fmt.Sprintf("cellphone=$%d", argID))
		args = append(args, *input.CellPhone)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE persons SET %s WHERE id=%d", setQuery, id)

	p.db.Exec(query, args...)
}

func (p Persons) Delete(id int) {
	//p.db.Query(fmt.Sprintf("DELETE FROM persons WHERE id = %d", id))
	p.db.Exec("DELETE FROM persons WHERE id = $1 ", id)
}
