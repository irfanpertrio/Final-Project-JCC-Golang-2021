package Keranjang

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
	table          = "Keranjang"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll Keranjang
func GetAll(ctx context.Context) ([]models.Keranjang, error) {

	var Keranjangs []models.Keranjang

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By Keranjang_id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var Keranjang models.Keranjang

		if err = rowQuery.Scan(&Keranjang.Keranjang_id,
			&Keranjang.Nama_barang,
			&Keranjang.Jumlah_barang,
			&Keranjang.Id_barang,
			&Keranjang.Catatan_pemesan,
			&Keranjang.Users_id); err != nil {
			return nil, err
		}

		Keranjangs = append(Keranjangs, Keranjang)
	}

	return Keranjangs, nil
}

// Insert Keranjang
func Insert(ctx context.Context, Keranjang models.Keranjang) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (keranjang_id, nama_barang, jumlah_barang, id_barang, catatan_pemesan, users_id) values('%v','%v','%v','%v','%v','%v')", table,
		Keranjang.Keranjang_id,
		Keranjang.Nama_barang,
		Keranjang.Jumlah_barang,
		Keranjang.Id_barang,
		Keranjang.Catatan_pemesan,
		Keranjang.Users_id,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Keranjang
func Update(ctx context.Context, Keranjang models.Keranjang, id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set keranjang_id ='%d', nama_barang ='%s', jumlah_barang ='%d', id_barang ='%s', catatan_pemesan ='%s', users_id ='%d' where keranjang_id = %s",
		table,
		Keranjang.Keranjang_id,
		Keranjang.Nama_barang,
		Keranjang.Jumlah_barang,
		Keranjang.Id_barang,
		Keranjang.Catatan_pemesan,
		Keranjang.Users_id,
		id,
	)

	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Keranjang
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where keranjang_id = %s", table, id)

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
