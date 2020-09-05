package models

type RequestSync struct {
	Mode         string `json:"mode" form:"mode"`
	Model        string `json:"model" form:"model"`
	username     string `json:"username" form:"username"`
	DeviceId     string `json:"devid" form:"devid"`
	Varian       string `json:"varian"`
	ItemId       int64  `json:"itemid"`
	PerusahaanNo string `json:"perusahaanno"`
	Data         string
}

type ResponseSync struct {
	Status  bool        `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}
