package main

import (
	"api-mysql/biodata_diri"
	"api-mysql/daftar_alamat"
	"api-mysql/keranjang"
	"api-mysql/models"
	"api-mysql/pembayaran"
	"api-mysql/pesanan"
	"api-mysql/ulasan"
	"api-mysql/users"
	"api-mysql/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := httprouter.New()

	router.GET("/users", BasicAuth(GetUsers))
	router.POST("/users/create", PostUsers)
	router.PUT("/users/:id/update", UpdateUsers)
	router.DELETE("/users/:id/delete", BasicAuth(DeleteUsers))

	router.GET("/pembayaran", BasicAuth(GetPembayaran))
	router.POST("/pembayaran/create", BasicAuth(PostPembayaran))
	router.PUT("/pembayaran/:id/update", BasicAuth(UpdatePembayaran))
	router.DELETE("/pembayaran/:id/delete", BasicAuth(DeletePembayaran))

	router.GET("/pesanan", BasicAuth(GetPesanan))
	router.POST("/pesanan/create", BasicAuth(PostPesanan))
	router.PUT("/pesanan/:id/update", BasicAuth(UpdatePesanan))
	router.DELETE("/pesanan/:id/delete", BasicAuth(DeletePesanan))

	router.GET("/keranjang", BasicAuth(GetKeranjang))
	router.POST("/keranjang/create", BasicAuth(PostKeranjang))
	router.PUT("/keranjang/:id/update", BasicAuth(UpdateKeranjang))
	router.DELETE("/keranjang/:id/delete", BasicAuth(DeleteKeranjang))

	router.GET("/biodata_diri", BasicAuth(GetBiodata_diri))
	router.POST("/biodata_diri/create", BasicAuth(PostBiodata_diri))
	router.PUT("/biodata_diri/:id/update", BasicAuth(UpdateBiodata_diri))
	router.DELETE("/biodata_diri/:id/delete", BasicAuth(DeleteBiodata_diri))

	router.GET("/daftar_alamat", BasicAuth(GetDaftar_alamat))
	router.POST("/daftar_alamat/create", BasicAuth(PostDaftar_alamat))
	router.PUT("/daftar_alamat/:id/update", BasicAuth(UpdateDaftar_alamat))
	router.DELETE("/daftar_alamat/:id/delete", BasicAuth(DeleteDaftar_alamat))

	router.GET("/ulasan", BasicAuth(GetUlasan))
	router.POST("/ulasan/create", BasicAuth(PostUlasan))
	router.PUT("/ulasan/:id/update", BasicAuth(UpdateUlasan))
	router.DELETE("/ulasan/:id/delete", BasicAuth(DeleteUlasan))

	// untuk menampilkan file html di folder public
	router.NotFound = http.FileServer(http.Dir("public"))

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

//basicauthentication
func BasicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == "pengguna" && password == "sayhai" {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

// Read
// GetBiodata_diri
func GetUsers(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	Userss, err := users.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, Userss, http.StatusOK)
}

// Create
// Postusers
func PostUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var bio models.Users

	if err := json.NewDecoder(r.Body).Decode(&bio); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := users.Insert(ctx, bio); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Updateusers
func UpdateUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var bio models.Users

	if err := json.NewDecoder(r.Body).Decode(&bio); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idusers = ps.ByName("id")

	if err := users.Update(ctx, bio, idusers); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
func DeleteUsers(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idusers = ps.ByName("id")

	if err := users.Delete(ctx, idusers); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

//////////////////////////////////////////////////////////

// Read
// GetBiodata_diri
func GetPembayaran(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	Pembayarans, err := pembayaran.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, Pembayarans, http.StatusOK)
}

// Create
// Postusers
func PostPembayaran(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var bio models.Pembayaran

	if err := json.NewDecoder(r.Body).Decode(&bio); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := pembayaran.Insert(ctx, bio); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Updateusers
func UpdatePembayaran(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var bayar models.Pembayaran

	if err := json.NewDecoder(r.Body).Decode(&bayar); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idPembayaran = ps.ByName("id")

	if err := pembayaran.Update(ctx, bayar, idPembayaran); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeletePembayaran
func DeletePembayaran(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idPembayaran = ps.ByName("id")

	if err := pembayaran.Delete(ctx, idPembayaran); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

//////////////////////////////////////////////////////////

// Read
// GetBiodata_diri
func GetPesanan(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	Pesanans, err := pesanan.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, Pesanans, http.StatusOK)
}

// Create
// Postusers
func PostPesanan(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var bio models.Pesanan

	if err := json.NewDecoder(r.Body).Decode(&bio); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := pesanan.Insert(ctx, bio); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Updateusers
func UpdatePesanan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var bayar models.Pesanan

	if err := json.NewDecoder(r.Body).Decode(&bayar); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idPesanan = ps.ByName("id")

	if err := pesanan.Update(ctx, bayar, idPesanan); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeletePembayaran
func DeletePesanan(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idPesanan = ps.ByName("id")

	if err := pesanan.Delete(ctx, idPesanan); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

//////////////////////////////////////////////////////////////
// Read
// GetBiodata_diri
func GetKeranjang(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	Keranjangs, err := keranjang.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, Keranjangs, http.StatusOK)
}

// Create
// Postusers
func PostKeranjang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var bio models.Keranjang

	if err := json.NewDecoder(r.Body).Decode(&bio); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := keranjang.Insert(ctx, bio); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Updateusers
func UpdateKeranjang(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var bayar models.Keranjang

	if err := json.NewDecoder(r.Body).Decode(&bayar); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idKeranjang = ps.ByName("id")

	if err := keranjang.Update(ctx, bayar, idKeranjang); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeletePembayaran
func DeleteKeranjang(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idKeranjang = ps.ByName("id")

	if err := keranjang.Delete(ctx, idKeranjang); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////
// Read
// GetBiodata_diri
func GetDaftar_alamat(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	Daftar_alamats, err := daftar_alamat.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, Daftar_alamats, http.StatusOK)
}

// Create
// Postusers
func PostDaftar_alamat(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var bio models.Daftar_alamat

	if err := json.NewDecoder(r.Body).Decode(&bio); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := daftar_alamat.Insert(ctx, bio); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Updateusers
func UpdateDaftar_alamat(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var loc models.Daftar_alamat

	if err := json.NewDecoder(r.Body).Decode(&loc); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idDaftar_alamat = ps.ByName("id")

	if err := daftar_alamat.Update(ctx, loc, idDaftar_alamat); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeletePembayaran
func DeleteDaftar_alamat(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idDaftar_alamat = ps.ByName("id")

	if err := daftar_alamat.Delete(ctx, idDaftar_alamat); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////
// Read
// Getbiodata_diri
func GetBiodata_diri(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	Biodata_diris, err := biodata_diri.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, Biodata_diris, http.StatusOK)
}

// Create
// PostBiodata_diri
func PostBiodata_diri(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var bio models.Biodata_diri

	if err := json.NewDecoder(r.Body).Decode(&bio); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := biodata_diri.Insert(ctx, bio); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateBiodata_diri
func UpdateBiodata_diri(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var bio models.Biodata_diri

	if err := json.NewDecoder(r.Body).Decode(&bio); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idBiodata_diri = ps.ByName("id")

	if err := biodata_diri.Update(ctx, bio, idBiodata_diri); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteBiodata_diri
func DeleteBiodata_diri(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idBiodata_diri = ps.ByName("id")

	if err := biodata_diri.Delete(ctx, idBiodata_diri); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////
// Read
// Getbiodata_diri
func GetUlasan(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	Ulasans, err := ulasan.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, Ulasans, http.StatusOK)
}

// Create
// PostBiodata_diri
func PostUlasan(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var cek models.Ulasan

	if err := json.NewDecoder(r.Body).Decode(&cek); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := ulasan.Insert(ctx, cek); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateBiodata_diri
func UpdateUlasan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var bio models.Ulasan

	if err := json.NewDecoder(r.Body).Decode(&bio); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idUlasan = ps.ByName("id")

	if err := ulasan.Update(ctx, bio, idUlasan); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteBiodata_diri
func DeleteUlasan(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idUlasan = ps.ByName("id")

	if err := ulasan.Delete(ctx, idUlasan); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}
