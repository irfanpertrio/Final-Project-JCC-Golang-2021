package users

import (
	"api-mysql/config"
	"api-mysql/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const (
	table          = "users"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll Users
func GetAll(ctx context.Context) ([]models.Users, error) {

	var Userss []models.Users

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By users_id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var Users models.Users

		if err = rowQuery.Scan(&Users.Users_id,
			&Users.Username,
			&Users.Password,
			&Users.Role); err != nil {
			return nil, err
		}

		Userss = append(Userss, Users)
	}

	return Userss, nil
}

// Insert Users
func Insert(ctx context.Context, Users models.Users) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (users_id, username, password, role) values('%v','%v','%v','%v')", table,
		Users.Users_id,
		Users.Username,
		Users.Password,
		Users.Role,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Users
func Update(ctx context.Context, Users models.Users, id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set user_id ='%d', username ='%s', password ='%s', role ='%s' where users_id = %s",
		table,
		Users.Users_id,
		Users.Username,
		Users.Password,
		Users.Role,
		id,
	)

	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Users
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where users_id = %s", table, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
