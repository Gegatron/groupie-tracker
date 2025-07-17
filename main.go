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
	   
	Locations []string `json:"locations"`
	
}
type DatesData struct {
	
	
	Dates     []string   `json:"dates"`
}
type RelationsData struct {
	
	Relations map[string][]string `json:"datesLocations"`
	
}
var LocationResponse struct {
	Index []LocationData `json:"index"`
}
var DatesResponse struct {
	Index []DatesData `json:"index"`
}
var RelationResponse struct {
	Index []RelationsData `json:"index"`
}
func main() {
	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))
	fetchdata()
	
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
		"Artist" : artists[id-1],
		"Locations":LocationResponse.Index[id-1],
		"Dates" : DatesResponse.Index[id-1],
		"Relations" : RelationResponse.Index[id-1],
	})
}
func fetchdata(){
artist, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return
	}
	err = json.NewDecoder(artist.Body).Decode(&artists)
	if err != nil {
		return
	}
	locations, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Println("fffc")
		return
	}
	err = json.NewDecoder(locations.Body).Decode(&LocationResponse)
	if err != nil {
		fmt.Println("ffbf")
		return
	}
	dates, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		fmt.Println("ffgf")
		return
	}
	err = json.NewDecoder(dates.Body).Decode(&DatesResponse)
	if err != nil {
		fmt.Println("ffaf",err)
		return
	}
	
	relations, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("ffef")
		return
	}
	err = json.NewDecoder(relations.Body).Decode(&RelationResponse)
	if err != nil {
		fmt.Println("ffpf",err)
		return
	}
	

}