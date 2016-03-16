package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"./model"
	"net/http"
	"html/template"
)

func main() {

	http.Handle("/", func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("postgres", "postgres://postgres:123456@localhost/stock?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		rows, err := db.Query(`select s.id, b.name,  b.industry,ths_percent  from stock s join stock_base b on s.id=b.id
			where icf_level=50 and ths_percent>=90  and c_date=current_date order by s.ths_percent desc limit 10`)

		if err != nil {
			log.Fatal(err)
		}

		arr := []model.Stock{}
		for rows.Next() {
			s := model.Stock{}
			err = rows.Scan(&s.Id, &s.Name, &s.Industry, &s.Niucha_percent)
			s.WebId = s.Id[2:]
			s.WebId += "." + s.Id[0:2]
			arr = append(arr, s)
		}
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, arr)
	})

	http.ListenAndServe(":80", nil)
}
