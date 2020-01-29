package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"

	"log"

	"github.com/gorilla/mux"
	"swapid/entity"
	"swapid/global"
	"swapid/utils"
)

func main() {
	var err error
	global.DB, err = gorm.Open(global.DBDialect, global.DBName)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	defer global.DB.Close()

	log.Printf("swapid is going to listen on %s\n", global.ListeningAddress)

	// updateSchema()

	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter()

	// CRUD
	router.HandleFunc("/characters", createCharacter).Methods("POST")
	router.HandleFunc("/characters/{id}", readCharacter).Methods("GET")
	router.HandleFunc("/characters/{id}", updateCharacter).Methods("PUT")
	router.HandleFunc("/characters/{id}", deleteCharacter).Methods("DELETE")

	router.HandleFunc("/characters", listCharacters).Methods("GET")

	log.Fatal(http.ListenAndServe(global.ListeningAddress, router))
}

func readCharacter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var character entity.Character
	if err := global.DB.Where("id=?", id).Find(&character).Error; err != nil {
		utils.HandleError(w, fmt.Sprintf("DB error while reading character with id: %s", id))
		return
	}

	utils.SetContentType(w, utils.JSON{})
	json.NewEncoder(w).Encode(character)
	log.Printf("Successfully read the cheracter with id %d (%s)", character.ID, character.Name)
}

func listCharacters(w http.ResponseWriter, r *http.Request) {

	var characters []entity.Character
	if err := global.DB.Find(&characters).Error; err != nil {
		utils.HandleError(w, fmt.Sprintf("DB error while reading characters"))
		return
	}

	utils.SetContentType(w, utils.JSON{})
	json.NewEncoder(w).Encode(characters)
	log.Printf("Successfully read the cheracters")
}

func deleteCharacter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var character entity.Character
	if err := global.DB.Where("id=?", id).Find(&character).Error; err != nil {
		utils.HandleError(w, fmt.Sprintf("DB error while reading character with id: %s %v", id, err))
		return
	}

	if err := global.DB.Delete(&character).Error; err != nil {
		utils.HandleError(w, fmt.Sprintf("DB error while delete character with id: %s %v", id, err))
		return
	}

	utils.SetContentType(w, utils.JSON{})
	json.NewEncoder(w).Encode(character)

	log.Printf("Successfully delete the cheracter with id %d (%s)", character.ID, character.Name)

}

func updateCharacter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var updateCharacter entity.Character

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&updateCharacter); err != nil {
		utils.HandleError(w, fmt.Sprintf("JSON decoding error %v\n", err))
		return
	}

	var character entity.Character
	if err := global.DB.Where("id=?", id).Find(&character).Error; err != nil {
		utils.HandleError(w, fmt.Sprintf("DB error while reading character with id: %s", id))
		return
	}

	character.Name = updateCharacter.Name

	if err := global.DB.Save(&character); err != nil {
		utils.HandleError(w, fmt.Sprintf("DB error while reading character with id: %v", &character))
		return
	}

	utils.SetContentType(w, utils.JSON{})
	json.NewEncoder(w).Encode(character)
	log.Printf("Successfully update the character with id %d (%s)", character.ID, character.Name)

}

func createCharacter(w http.ResponseWriter, r *http.Request) {
	var character entity.Character

	err := json.NewDecoder(r.Body).Decode(&character)
	if err != nil {
		utils.HandleError(w, fmt.Sprintf("JSON encoding error %s", err.Error()))
		return
	}

	if err = global.DB.Create(&character).Error; err != nil {
		utils.HandleError(w, fmt.Sprintf("DB error while creating character %s\n", err.Error()))
		return
	}

	utils.SetContentType(w, utils.JSON{})
	json.NewEncoder(w).Encode(character)
	log.Printf("Successfully created character with ID %d\n", character.ID)

}

func updateSchema() {
	global.DB.AutoMigrate(&entity.Character{})
}
