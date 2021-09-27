package main

import (
	"api-mysql/biodata_diri"
	Biodata_diri "api-mysql/biodata_diri"
	"api-mysql/models"
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

	// router.GET("/users", Basic_auth(GetUsers))
	// router.POST("/users/create", PostUsers)
	// router.PUT("/users/:id/update", UpdateUsers)
	// router.DELETE("/users/:id/delete", Basic_auth(DeleteUsers))

	// router.GET("/pembayaran", Basic_auth(GetPembayaran))
	// router.POST("/pembayaran/create", Basic_auth(PostPembayaran))
	// router.PUT("/pembayaran/:id/update", Basic_auth(UpdatePembayaran))
	// router.DELETE("/pembayaran/:id/delete", Basic_auth(DeletePembayaran))

	// router.GET("/pesanan", Basic_auth(GetPesanan))
	// router.POST("/pesanan/create", Basic_auth(PostPesanan))
	// router.PUT("/pesanan/:id/update", Basic_auth(UpdatePesanan))
	// router.DELETE("/pesanan/:id/delete", Basic_auth(DeletePesanan))

	// router.GET("/keranjang", Basic_auth(GetKeranjang))
	// router.POST("/keranjang/create", Basic_auth(PostKeranjang))
	// router.PUT("/keranjang/:id/update", Basic_auth(UpdateKeranjang))
	// router.DELETE("/keranjang/:id/delete", Basic_auth(DeleteKeranjang))

	router.GET("/biodata_diri", Basic_auth(GetBiodata_diri))
	router.POST("/biodata_diri/create", Basic_auth(PostBiodata_diri))
	router.PUT("/biodata_diri/:id/update", Basic_auth(UpdateBiodata_diri))
	router.DELETE("/biodata_diri/:id/delete", Basic_auth(DeleteBiodata_diri))

	// router.GET("/daftar_alamat", Basic_auth(GetDaftar_alamat))
	// router.POST("/daftar_alamat/create", Basic_auth(PostDaftar_alamat))
	// router.PUT("/daftar_alamat/:id/update", Basic_auth(UpdateDaftar_alamat))
	// router.DELETE("/daftar_alamat/:id/delete", Basic_auth(DeleteDaftar_alamat))

	// router.GET("/ulasan", Basic_auth(GetUlasan))
	// router.POST("/ulasan/create", Basic_auth(PostUlasan))
	// router.PUT("/ulasan/:id/update", Basic_auth(UpdateUlasan))
	// router.DELETE("/ulasan/:id/delete", Basic_auth(DeleteUlasan))

	// untuk menampilkan file html di folder public
	router.NotFound = http.FileServer(http.Dir("public"))

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

//basicauthentication
func Basic_auth(h httprouter.Handle) httprouter.Handle {
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

	if err := Biodata_diri.Insert(ctx, bio); err != nil {
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

	if err := Biodata_diri.Update(ctx, bio, idBiodata_diri); err != nil {
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

	if err := Biodata_diri.Delete(ctx, idBiodata_diri); err != nil {
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

////////////////////////////////////////////////////////////////////
