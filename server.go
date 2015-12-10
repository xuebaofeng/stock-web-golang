package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/gzip"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"github.com/xuebaofeng/stock-web-golang/model"
)

func main() {
	m := martini.Classic()
	m.Use(gzip.All())
	m.Use(martini.Static("public"))

	m.Use(render.Renderer(render.Options{
		Extensions: []string{".html"}, // Specify extensions to load for templates.
		Charset: "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
	}))

	m.Get("/", func(r render.Render) {
		db, err := sql.Open("postgres", "postgres://postgres:123456@localhost/stock?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		rows, err := db.Query(`select s.id, b.name,  b.industry,ths_percent  from stock s join stock_base b on s.id=b.id
			where icf_level=50 and ths_percent>=80  and c_date=current_date order by s.ths_percent desc;`)

		if err != nil {
			log.Fatal(err)
		}

		arr := []model.Stock{}
		for rows.Next() {
			s := model.Stock{}
			err = rows.Scan(&s.Id, &s.Name, &s.Industry, &s.Niucha_percent)
			s.ShortId = s.Id[2:]
			arr = append(arr, s)
		}
		r.HTML(200, "index", arr)
	})


	m.RunOnAddr(":80")
}
