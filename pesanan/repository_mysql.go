package Pesanan

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
	table          = "Pesanan"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll Pesanan
func GetAll(ctx context.Context) ([]models.Pesanan, error) {

	var Pesanans []models.Pesanan

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By Pesanan_id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var Pesanan models.Pesanan

		if err = rowQuery.Scan(&Pesanan.Id_barang,
			&Pesanan.Jumlah_barang,
			&Pesanan.Alamat_id,
			&Pesanan.Keranjang_id); err != nil {
			return nil, err
		}

		Pesanans = append(Pesanans, Pesanan)
	}

	return Pesanans, nil
}

// Insert Pesanan
func Insert(ctx context.Context, Pesanan models.Pesanan) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (id_barang, jumlah_barang, alamat_id, keranjang_id) values('%v','%v','%v','%v')", table,
		Pesanan.Id_barang,
		Pesanan.Jumlah_barang,
		Pesanan.Alamat_id,
		Pesanan.Keranjang_id,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Pesanan
func Update(ctx context.Context, Pesanan models.Pesanan, id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set id_barang ='%d', jumlah_barang ='%d', alamat_id ='%d', keranjang_id ='%d' where pesanan_id = %s",
		table,
		Pesanan.Id_barang,
		Pesanan.Jumlah_barang,
		Pesanan.Alamat_id,
		Pesanan.Keranjang_id,
		id,
	)

	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Pesanan
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where pesanan_id = %s", table, id)

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
