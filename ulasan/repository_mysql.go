package Ulasan

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
	table          = "Ulasan"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll Ulasan
func GetAll(ctx context.Context) ([]models.Ulasan, error) {

	var Ulasans []models.Ulasan

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By Ulasan_id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var Ulasan models.Ulasan

		if err = rowQuery.Scan(&Ulasan.Ulasan_id,
			&Ulasan.Review,
			&Ulasan.Users_id,
			&Ulasan.Keranjang_id); err != nil {
			return nil, err
		}

		Ulasans = append(Ulasans, Ulasan)
	}

	return Ulasans, nil
}

// Insert Ulasan
func Insert(ctx context.Context, Ulasan models.Ulasan) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (ulasan_id, review, users_id, keranjang_id) values('%v','%v','%v','%v')", table,
		Ulasan.Ulasan_id,
		Ulasan.Review,
		Ulasan.Users_id,
		Ulasan.Keranjang_id,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Ulasan
func Update(ctx context.Context, Ulasan models.Ulasan, id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set ulasan_id ='%d', review ='%s', users_id ='%d', keranjang_id ='%d' where ulasan_id = %s",
		table,
		Ulasan.Ulasan_id,
		Ulasan.Review,
		Ulasan.Users_id,
		Ulasan.Keranjang_id,
		id,
	)

	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Ulasan
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where ulasan_id = %s", table, id)

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
