package main

import "net/http"
import "view"
import "model"

import (
_ "github.com/go-sql-driver/mysql"
)

func main(){
	http.HandleFunc("/", view.IndexHandler)
	http.HandleFunc("/tree", view.TreeHandler)
	http.HandleFunc("/gen_token", view.GenHashHandler)
	http.ListenAndServe(":8080", nil)
	defer model.DB.Close()
}
