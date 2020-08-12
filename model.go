package main

type Kelas struct {
	Id     int    `form:"id" json:"id"`
	Nama   string `form:"nama" json:"nama"`
	Active string `form:"active" json:"active"`
}

type Response struct {
	Status      int     `json:"status"`
	Message     string  `json:"message"`
	TotalPage   int     `json:"totalPage"`
	CurrentPage int     `json:"currentPage"`
	Data        []Kelas `json:"data"`
}
