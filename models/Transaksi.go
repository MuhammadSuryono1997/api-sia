package models

type Transaksi struct {
	IdTransaksi  int64  `json:"id_transaksi"`
	IdUser       int64  `json:"id_user"`
	NoReff       int64  `json:"no_reff"`
	Keterangan   string `json:"keterangan"`
	Lampiran     string `json:"lampiran"`
	TglInput     string `json:"tanggal_input"`
	TglTransaksi string `json:"tanggal_transaksi"`
	JenisSaldo   string `json:"jenis_saldo"`
	Saldo        int32  `json:"saldo"`
}
