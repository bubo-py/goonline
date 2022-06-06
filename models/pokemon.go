package models

// Pokemon is a struct for actual database pokemon representation
type Pokemon struct {
	ID           int
	NameEnglish  string `json:"english"`
	NameJapanese string `json:"japanese"`
	Type         string `json:"type"`
	Hp           uint   `json:"hp"`
	Attack       uint   `json:"attack"`
	Defense      uint   `json:"defense"`
	Sp_attack    uint   `json:"sp. attack"`
	Sp_defense   uint   `json:"sp. defense"`
	Speed        uint   `json:"speed"`
}

// JsonPokemon is a struct for retrieval data from provided json
type JsonPokemon struct {
	ID   int
	Name struct {
		English  string `json:"english"`
		Japanese string `json:"japanese"`
	} `json:"name"`

	Type []string `json:"type"`
	Base struct {
		Hp         uint `json:"hp"`
		Attack     uint `json:"attack"`
		Defense    uint `json:"defense"`
		Sp_attack  uint `json:"sp. attack"`
		Sp_defense uint `json:"sp. defense"`
		Speed      uint `json:"speed"`
	} `json:"base"`
}
