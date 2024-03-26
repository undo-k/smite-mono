package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	output, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(output)
	if err != nil {
		return err
	}

	return nil
}

type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	minutes := d.Duration / time.Minute
	seconds := (d.Duration % time.Minute) / time.Second
	return json.Marshal(fmt.Sprintf("%02d:%02d", minutes, seconds))
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}

// func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) error {
// 	statusCode := http.StatusBadRequest

// 	if len(status) > 0 {
// 		statusCode = status[0]
// 	}

// 	var payload jsonResponse
// 	payload.Error = true
// 	payload.Message = err.Error()

// 	return app.writeJSON(w, statusCode, payload)
// }

func Gods() []string {
	return []string{
		"Achilles",
		"Agni",
		"Ah Muzen Cab",
		"Ah Puch",
		"Amaterasu",
		"Anhur",
		"Anubis",
		"Ao Kuang",
		"Aphrodite",
		"Apollo",
		"Arachne",
		"Ares",
		"Artemis",
		"Artio",
		"Athena",
		"Atlas",
		"Awilix",
		"Baba Yaga",
		"Bacchus",
		"Bakasura",
		"Bake Kujira",
		"Baron Samedi",
		"Bastet",
		"Bellona",
		"Cabrakan",
		"Camazotz",
		"Cerberus",
		"Cernunnos",
		"Chaac",
		"Change",
		"Charon",
		"Charybdis",
		"Chernobog",
		"Chiron",
		"Chronos",
		"Cliodhna",
		"Cthulhu",
		"Cu Chulainn",
		"Cupid",
		"Da Ji",
		"Danzaburou",
		"Discordia",
		"Erlang Shen",
		"Eset",
		"Fafnir",
		"Fenrir",
		"Freya",
		"Ganesha",
		"Geb",
		"Gilgamesh",
		"Guan Yu",
		"Hachiman",
		"Hades",
		"He Bo",
		"Heimdallr",
		"Hel",
		"Hera",
		"Hercules",
		"Horus",
		"Hou Yi",
		"Hun Batz",
		"Ishtar",
		"Ix Chel",
		"Izanami",
		"Janus",
		"Jing Wei",
		"Jormungandr",
		"Kali",
		"Khepri",
		"King Arthur",
		"Kukulkan",
		"Kumbhakarna",
		"Kuzenbo",
		"Lancelot",
		"Loki",
		"Maman Brigitte",
		"Martichoras",
		"Maui",
		"Medusa",
		"Mercury",
		"Merlin",
		"Morgan Le Fay",
		"Mulan",
		"Ne Zha",
		"Neith",
		"Nemesis",
		"Nike",
		"Nox",
		"Nu Wa",
		"Nut",
		"Odin",
		"Olorun",
		"Osiris",
		"Pele",
		"Persephone",
		"Poseidon",
		"Ra",
		"Raijin",
		"Rama",
		"Ratatoskr",
		"Ravana",
		"Scylla",
		"Serqet",
		"Set",
		"Shiva",
		"Skadi",
		"Sobek",
		"Sol",
		"Sun Wukong",
		"Surtr",
		"Susano",
		"Sylvanus",
		"Terra",
		"Thanatos",
		"The Morrigan",
		"Thor",
		"Thoth",
		"Tiamat",
		"Tsukuyomi",
		"Tyr",
		"Ullr",
		"Vamana",
		"Vulcan",
		"Xbalanque",
		"Xing Tian",
		"Yemoja",
		"Ymir",
		"Yu Huang",
		"Zeus",
		"Zhong Kui",
	}
}
