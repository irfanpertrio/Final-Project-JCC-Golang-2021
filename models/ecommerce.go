package models

type (
	Users struct {
		Users_id int    `json:"users_id"`
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	Pembayaran struct {
		Pembayaran_id int `json:"pembayaran_id"`
		Kartu_kredit  int `json:"kartu_kredit"`
		Kredivo       int `json:"kredivo"`
		Debit         int `json:"debit"`
		Users_id      int `json:"users_id"`
	}

	Pesanan struct {
		Id_barang     int `json:"id_barang"`
		Jumlah_barang int `json:"jumlah_barang"`
		Alamat_id     int `json:"alamat_id"`
		Keranjang_id  int `json:"keranjang_id"`
	}

	Keranjang struct {
		Keranjang_id    int    `json:"keranjang_id"`
		Nama_barang     string `json:"nama_barang"`
		Jumlah_barang   int    `json:"jumlah_barang"`
		Id_barang       string `json:"id_barang"`
		Catatan_pemesan string `json:"Catatan_pemesan"`
		Users_id        int    `json:"users_id"`
	}

	Biodata_diri struct {
		Biodata_id    int    `json:"biodata_id"`
		Nama          string `json:"nama"`
		Tanggal_lahir string `json:"tanggal_lahir"`
		Jenis_kelamin string `json:"jenis_kelamin"`
		Email         string `json:"email"`
		No_hp         int    `json:"no_hp"`
		Profil_pic    string `json:"profil_pic"`
		Users_id      int    `json:"users_id"`
	}

	Daftar_alamat struct {
		Alamat_id      int    `json:"alamat_id"`
		Label_alamat   string `json:"label_alamat"`
		Nama_penerima  string `json:"nama_penerima"`
		No_hp          int    `json:"no_hp"`
		Kota_kecamatan string `json:"kota_kecamatan"`
		Kode_pos       int    `json:"kode_pos"`
		Alamat         string `json:"alamat"`
		Users_id       int    `json:"users_id"`
	}

	Ulasan struct {
		Ulasan_id    int    `json:"ulasan_id"`
		Review       string `json:"review"`
		Users_id     int    `json:"users_id"`
		Keranjang_id int    `json:"keranjang_id"`
	}
)
