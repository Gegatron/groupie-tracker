package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

var artists []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	LocationsURL string   `json:"locations"`
	DatesURL     string   `json:"concertDates"`
	RelationsURL string   `json:"relations"`
}


func main() {
	artist, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return
	}
	err = json.NewDecoder(artist.Body).Decode(&artists)
	if err != nil {
		return
	}
	http.HandleFunc("/", tracking)
	http.ListenAndServe(":8080", nil)
}

func tracking(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println("error")
		return
	}
	temp.Execute(w, artists)
}
