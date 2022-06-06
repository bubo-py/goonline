package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/bubo-py/goonline/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database variables
var Instance *gorm.DB
var dbError error

// Connect opens the connection to MySQL database
func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to the database")
	}
	log.Println("Connected to the database!")
}

// Migrate creates fields in database from models
func Migrate() {
	Instance.AutoMigrate(&models.Profile{})
	Instance.AutoMigrate(&models.Pokemon{})
	log.Println("Database migration completed!")
}

// InitializeData initializes data from the go-test-db.json file
func InitializeData() {
	// open json file
	jsonFile, err := os.Open("go-test-db.json")
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()
	byteJson, _ := ioutil.ReadAll(jsonFile)

	var jsonPokemon []models.JsonPokemon
	var pokemon models.Pokemon

	json.Unmarshal(byteJson, &jsonPokemon)

	// map json structure to the GORM model structure, then create the record
	for i := 0; i <= len(jsonPokemon)-1; i++ {
		pokemon.ID = jsonPokemon[i].ID
		pokemon.NameEnglish = jsonPokemon[i].Name.English
		pokemon.NameJapanese = jsonPokemon[i].Name.Japanese
		pokemon.Type = concatTypes(jsonPokemon[i].Type)
		pokemon.Hp = jsonPokemon[i].Base.Hp
		pokemon.Attack = jsonPokemon[i].Base.Attack
		pokemon.Defense = jsonPokemon[i].Base.Defense
		pokemon.Sp_attack = jsonPokemon[i].Base.Sp_attack
		pokemon.Sp_defense = jsonPokemon[i].Base.Sp_defense
		pokemon.Speed = jsonPokemon[i].Base.Speed

		Instance.Create(&pokemon)
	}

}

// concatTypes helps to concat types if there's more than one of them
func concatTypes(types []string) string {
	return strings.Join(types, ", ")
}
