package pembayaran

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
	table          = "pembayaran"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll Pembayaran
func GetAll(ctx context.Context) ([]models.Pembayaran, error) {

	var Pembayarans []models.Pembayaran

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By pembayaran_id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var Pembayaran models.Pembayaran

		if err = rowQuery.Scan(&Pembayaran.Pembayaran_id,
			&Pembayaran.Kartu_kredit,
			&Pembayaran.Kredivo,
			&Pembayaran.Debit,
			&Pembayaran.Users_id); err != nil {
			return nil, err
		}

		Pembayarans = append(Pembayarans, Pembayaran)
	}

	return Pembayarans, nil
}

// Insert Pembayaran
func Insert(ctx context.Context, Pembayaran models.Pembayaran) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (pembayaran_id, users_id, kartu_kredit, kredivo, debit) values('%v','%v','%v','%v','%v')", table,
		Pembayaran.Pembayaran_id,
		Pembayaran.Users_id,
		Pembayaran.Kartu_kredit,
		Pembayaran.Kredivo,
		Pembayaran.Debit,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Pembayaran
func Update(ctx context.Context, Pembayaran models.Pembayaran, id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set user_id ='%d', kartu_kredit ='%d', kredivo ='%d', debit ='%d' where pembayaran_id = %s",
		table,
		Pembayaran.Users_id,
		Pembayaran.Kartu_kredit,
		Pembayaran.Kredivo,
		Pembayaran.Debit,
		id,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Pembayaran
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where pembayaran_id = %s", table, id)

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
