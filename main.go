package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func init() {
	dsn := "host=CONB20BB005_2 port=5432 user=b20bb005 password=toor dbname=userdb sslmode=disable"
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Println(err)

		panic("Failed to connect to database")
	}

	db.AutoMigrate(&User{})
}

func main() {
	http.HandleFunc("/users", usersHandler)
	http.ListenAndServe(":8080", nil)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var users []User
		db.Find(&users)
		json.NewEncoder(w).Encode(users)
	case "POST":
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		db.Create(&user)
		json.NewEncoder(w).Encode(user)
	case "PUT":
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		db.Save(&user)
		json.NewEncoder(w).Encode(user)
	case "DELETE":
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		db.Delete(&user)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
