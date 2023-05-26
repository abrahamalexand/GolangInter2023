package main

//ALEXANDER ABRAHAM PASARIBU
import (
	"alexander_50420116_pert4/handler" // Ganti dengan nama folder kalian masingmasing
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/", handler.API)
	log.Println("localhost : 8016")   //Ganti 2 digit akhir port dengan 2 digit akhir NPM anda
	http.ListenAndServe(":8016", nil) //Ganti 2 digit akhir port dengan 2 digit akhir NPM anda
}
