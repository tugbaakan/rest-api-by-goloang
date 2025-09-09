package models

import (
	"errors"
	"fmt"

	"example.com/restapi/db"
	"example.com/restapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (item User) Save() error {

	query := `
	INSERT INTO users(email, password)
	values(?,?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(item.Password)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(item.Email, hashedPassword)

	if err != nil {
		return err
	}

	return err

}

func GetAllUsers() ([]User, error) {

	query := `
	select * from users
	`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []User

	for rows.Next() {
		var item User

		err := rows.Scan(&item.ID, &item.Email, &item.Password)

		if err != nil {
			return nil, err
		}

		items = append(items, item)

	}

	return items, nil

}

func GetUserById(id int64) (*User, error) {

	query := "Select * from users where id = ?"

	row := db.DB.QueryRow(query, id)

	var item User

	err := row.Scan(&item.Email, &item.Password)
	if err != nil {
		return nil, err
	}

	return &item, nil

}

func (item User) UpdateUser() error {

	query := `
		UPDATE users
		SET password = ?
		WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(item.Password, item.ID)

	if err != nil {
		return err
	}

	return err

}

func (e User) DeleteUser() error {

	query := `
		delete from users
		WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	if err != nil {
		return err
	}

	return err

}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password from users where email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	fmt.Println("user", u)

	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return err

}
