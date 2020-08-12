package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getKelas(w http.ResponseWriter, r *http.Request) {
	var kelas Kelas
	var arr_kelas []Kelas
	var response Response

	db := connect()
	defer db.Close()
	// get dari params url
	vars := mux.Vars(r)
	totalDataPerPage, _ := strconv.Atoi(vars["perPage"])
	page, _ := strconv.Atoi(vars["page"])

	// hitung total row dari table kelas - kemudian isi variable total data
	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM kelas").Scan(&totalData)

	// hitung total page
	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	// validasi human error ketika melebihi total page atatu page < 1
	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	// tentukan index pertama pada setiap page
	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id,nama,active from kelas limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&kelas.Id, &kelas.Nama, &kelas.Active); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_kelas = append(arr_kelas, kelas)
		}
	}

	response.Status = 200
	response.Message = "success"
	response.Data = arr_kelas
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
