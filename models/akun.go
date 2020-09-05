package models

type Akun struct {
	NoReff     string `json:"no_reff"`
	IdUser     int64  `json:"id_user"`
	NamaReff   string `json:"nama_reff"`
	Keterangan string `json:"keterangan"`
}
