package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var artists []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type LocationData struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

var LocationResponse struct {
	Index []LocationData `json:"index"`
}

func main() {
	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))
	artist, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return
	}
	err = json.NewDecoder(artist.Body).Decode(&artists)
	if err != nil {
		return
	}
	new, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Println("fff")
		return
	}
	err = json.NewDecoder(new.Body).Decode(&LocationResponse)
	if err != nil {
		fmt.Println("fff")
		return
	}

	http.HandleFunc("/", Tracking)
	http.HandleFunc("/infos/", ArtistsInfo)
	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func Tracking(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println("error")
		return
	}
	temp.Execute(w, artists)
}

func ArtistsInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/infos/"))
	if err != nil {
		fmt.Println(err)
		return
	}
	temp, err := template.ParseFiles("templates/infos.html")
	if err != nil {
		fmt.Println("error")
		return
	}

	fmt.Println(artists[id-1].ID)
	fmt.Println(LocationResponse.Index[id-1].Locations)
	temp.Execute(w, map[string]interface{}{
		"Artist" : artists,
		"Locations":LocationResponse,
	})
}
