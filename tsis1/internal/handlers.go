package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type AnimeCharacter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Series string `json:"series"`
}

type Database struct {
	sync.RWMutex
	Characters []AnimeCharacter
}

var animeDB = Database{}

func (db *Database) LoadDataFromFile(filePath string) error {
	db.Lock()
	defer db.Unlock()

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileContent, &db.Characters)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetAnimeList() []AnimeCharacter {
	db.RLock()
	defer db.RUnlock()
	return db.Characters
}

func (db *Database) GetAnimeDetails(id int) (AnimeCharacter, bool) {
	db.RLock()
	defer db.RUnlock()

	for _, character := range db.Characters {
		if character.ID == id {
			return character, true
		}
	}

	return AnimeCharacter{}, false
}

filePath := "api/anime_data.json"

func GetAnimeList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Load data from the JSON file before fetching the list
	if err := animeDB.LoadDataFromFile(filepath); err != nil {
		http.Error(w, "Failed to load anime data", http.StatusInternalServerError)
		return
	}

	characters := animeDB.GetAnimeList()
	json.NewEncoder(w).Encode(characters)
}

func GetAnimeDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Load data from the JSON file before fetching details
	if err := animeDB.LoadDataFromFile(filepath); err != nil {
		http.Error(w, "Failed to load anime data", http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	characterID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid character ID", http.StatusBadRequest)
		return
	}

	character, found := animeDB.GetAnimeDetails(characterID)
	if found {
		json.NewEncoder(w).Encode(character)
	} else {
		http.Error(w, "Character not found", http.StatusNotFound)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("My Anime App\nAuthor: Dias"))
}
