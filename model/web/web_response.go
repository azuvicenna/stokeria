package web

type WebResponse struct {
	Code   int         `json:"code"`             // Contoh: 200, 400, 500
	Status string      `json:"status"`           // Contoh: "OK", "BAD REQUEST"
	Data   interface{} `json:"data,omitempty"`   // Data payload, hilang jika kosong (omitempty)
	Errors interface{} `json:"errors,omitempty"` // Detail error validasi, hilang jika sukses
}