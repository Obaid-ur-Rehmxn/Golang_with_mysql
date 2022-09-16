package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/insert", getData)
	http.ListenAndServe(":8000", nil)
	log.Print("Listening on port 8000")
}

func getData(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")
	log.Print(name, email, message)
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	query := "INSERT INTO feedback (name, email, message) VALUES ('" + name + "','" + email + "','" + message + "')"
	res, err := db.Exec(query)
	if err != nil {
		log.Println(err)
	} else {
		log.Print(res)
		w.Write([]byte("Success"))
	}

}
