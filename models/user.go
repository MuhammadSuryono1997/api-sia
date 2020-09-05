package models

type User struct {
	IdUser    int64  `json:"id_user"`
	Nama      string `json:"nama"`
	Jk        string `json:"jenis_kelamin"`
	Alamat    string `json:"alamat"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	LastLogin string `json:"last_login"`
}
