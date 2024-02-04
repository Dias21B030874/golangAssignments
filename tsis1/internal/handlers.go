package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"fmt"

	"github.com/gorilla/mux"
)

type AnimeCharacter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Series string `json:"series"`
}


var filePath = "api/anime_data.json"

func loadAnimeData() ([]AnimeCharacter, error) {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var characters []AnimeCharacter
	err = json.Unmarshal(fileContent, &characters)
	if err != nil {
		return nil, err
	}

	return characters, nil
}

func GetAnimeList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	characters, err := loadAnimeData()
	if err != nil {
		fmt.Println("Error loading anime data:", err)
		http.Error(w, "Failed to load anime data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(characters)
}

func GetAnimeDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	characterID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid character ID", http.StatusBadRequest)
		return
	}

	characters, err := loadAnimeData()
	if err != nil {
		fmt.Println("Error loading anime data:", err)
		http.Error(w, "Failed to load anime data", http.StatusInternalServerError)
		return
	}

	var foundCharacter AnimeCharacter
	found := false

	for _, character := range characters {
		if character.ID == characterID {
			foundCharacter = character
			found = true
			break
		}
	}

	if found {
		json.NewEncoder(w).Encode(foundCharacter)
	} else {
		http.Error(w, "Character not found", http.StatusNotFound)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("My Anime App\nAuthor: Dias"))
}
