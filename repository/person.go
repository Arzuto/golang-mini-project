package repository

import (
	"database/sql"
	"golang-mini-project/structs"
)

func GetAllPerson(db *sql.DB) ([]structs.Person, error) {
	sql := "SELECT * FROM PERSON"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []structs.Person 

	for rows.Next() {
		var person structs.Person

		err := rows.Scan(&person.ID, &person.FirstName, &person.LastName)
		if err != nil {
			return nil, err 
		}

		results = append(results, person)
	}

	err = rows.Err() 
	if err != nil {
		return nil, err 
	}

	return results, nil
}


func InsertPerson(db *sql.DB, person structs.Person) (err error) {
	sql := "INSERT INTO person (id, first_name, last_name) VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, person.ID, person.FirstName, person.LastName)

	return errs.Err()

}

func UpdatePerson(db *sql.DB, person structs.Person) (err error) {
	sql := "UPDATE person SET first_name = $1, last_name = $2 WHERE id = $3"

	errs := db.QueryRow(sql, person.FirstName, person.LastName, person.ID)

	return errs.Err()

}

func DeletePerson(db *sql.DB, person structs.Person) (err error) {
	sql := "DELETE FROM person WHERE id = $1"

	errs := db.QueryRow(sql, person.ID)

	return errs.Err()

}
