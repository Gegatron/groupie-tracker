package main

import (
	"encoding/json"

	"net/http"

	
)


var artists struct{
	ID int 
	Image string
	Name string
	Members []string
	CreationDate int
	FirstAlbum string
	locationsurl string
	datesurl string
	relationsurl string

}
var locations struct{
	ID int 
	locations []string
}
var dates struct{
	ID int
	dates []string
}
var relations struct{
	ID int
	relations string
}
func main(){
	
	artist,err:=http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err!=nil {
		return
	}
	err=json.NewDecoder(artist.Body).Decode(&artists)
	if err!=nil {
		return
	}

	
}
