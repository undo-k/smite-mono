package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/undo-k/smite-mono/protos/protos"
)

func (app *Config) GetGodByName(w http.ResponseWriter, r *http.Request) {
	godName := chi.URLParam(r, "godName")

	fmt.Printf("Fetching '%v' via gRPC\n", godName)

	godData, err := app.FetchGodViaGRPC(godName)

	if err != nil {

		err = fmt.Errorf("error fetching god via grpc: %w", err)
		fmt.Println(err)

		payload := jsonResponse{
			Error:   true,
			Message: "Could not retrieve god info",
		}
		app.writeJSON(w, http.StatusBadRequest, payload)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "",
		Data:    godData,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetGodList(w http.ResponseWriter, r *http.Request) {

	godList, err := app.FetchGodListViaGRPC()

	if err != nil {

		err = fmt.Errorf("error fetching godlist via grpc: %w", err)
		fmt.Println(err)

		payload := jsonResponse{
			Error:   true,
			Message: "Could not retrieve god list",
		}
		app.writeJSON(w, http.StatusInternalServerError, payload)
		return
	}

	app.writeJSON(w, http.StatusOK, jsonResponse{
		Error:   false,
		Message: "Success",
		Data:    godList.Gods,
	})
}

func (app *Config) PutGodByName(w http.ResponseWriter, r *http.Request) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	godName := chi.URLParam(r, "godName")

	god := protos.God{
		Id:         int32(rng.Intn(130)),
		Name:       godName,
		Role:       "gamer",
		AvgKills:   rng.Float32() * 8,
		AvgDeaths:  rng.Float32() * 15,
		AvgAssists: rng.Float32() * 19,
		AvgGold:    rng.Float32() * 43000,
		WinRate:    rng.Float32() * 99.0,
		HotItems:   []*protos.Item{},
	}

	god.HotItems = append(god.HotItems, &protos.Item{Id: 1, Name: "fist of destiny"})
	god.HotItems = append(god.HotItems, &protos.Item{Id: 1, Name: "kaboomerang"})
	god.HotItems = append(god.HotItems, &protos.Item{Id: 1, Name: "majin buu"})

	err := app.PutGodViaGRPC(&god)

	if err != nil {
		message := fmt.Errorf("could not insert god: %w", err)

		payload := jsonResponse{
			Error:   true,
			Message: message.Error(),
		}
		app.writeJSON(w, http.StatusBadRequest, payload)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Success",
	}

	app.writeJSON(w, http.StatusAccepted, payload)

}
func (app *Config) triggerAggregator(w http.ResponseWriter, r *http.Request) {
	number, _ := strconv.Atoi(chi.URLParam(r, "number"))

	err := app.triggerAggregatorViaGRPC(number)

	if err != nil {
		message := fmt.Errorf("could not trigger aggregator: %w", err)
		fmt.Println(err)

		payload := jsonResponse{
			Error:   true,
			Message: message.Error(),
		}
		app.writeJSON(w, http.StatusBadRequest, payload)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Success",
	}

	app.writeJSON(w, http.StatusAccepted, payload)

}
