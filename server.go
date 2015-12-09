package main

import (
	"github.com/go-martini/martini"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
)

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		db, err := sql.Open("postgres", "postgres://postgres:123456@localhost/stock?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		rows, err := db.Query(`select s.id, b.name,  b.industry,ths_percent  from stock s join stock_base b on s.id=b.id
			where icf_level=50 and ths_percent>=90  and c_date=current_date order by b.industry;`)

		if err != nil {
			log.Fatal(err)
		}

		var line string
		for rows.Next() {
			var id string
			var name string
			var industry string
			var ths_percent string
			err = rows.Scan(&id, &name, &industry, &ths_percent)
			s := fmt.Sprintf("%s\t%s\t%s\t%s\n", id, name, industry, ths_percent)
			fmt.Println(s)
			line += s

		}
		return line
	})
	m.RunOnAddr(":80")
}