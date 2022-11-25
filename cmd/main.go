package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id      int
	Model   string
	Company string
	Price   string
}

var sandbox_db *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := sandbox_db.Query("select * from sandbox_db.products;")
	defer rows.Close()

	if err != nil {
		log.Println(err)
	}
	products := []Product{}

	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
}
