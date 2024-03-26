package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

//developerId}/{signature}/{session}/{timestamp}/{queue}/date}/{hour}", app.GetMatchIdsByQueue)

//developerId}/{signature}/{session}/{timestamp}/{matchIds}", app.GetMatchDetailsBatch)

func (app *Config) GetMatchIdsByQueue(w http.ResponseWriter, r *http.Request) {
	devId := chi.URLParam(r, "developerId")
	signature := chi.URLParam(r, "signature") // md5_hash(devId + api method + authkey + timestamp)
	session := chi.URLParam(r, "session")
	timestamp := chi.URLParam(r, "timestamp") // "yyyyMMddHHmmss"
	queue := chi.URLParam(r, "queue")         // queue_id’s 435, 448, 445, 426, 451, 459, 450, & 440 are the only ones considered for player win/loss stats
	date := chi.URLParam(r, "date")           // yyyyMMdd
	hour := chi.URLParam(r, "hour")           // 0 - 23

	responseString := fmt.Sprintf("Received these url params: %s, %s, %s, %s, %s, %s, %s", devId, signature, session, timestamp, queue, date, hour)

	var response = struct {
		Error    bool     `json:"error"`
		Message  string   `json:"message"`
		MatchIds []string `json:"match_ids"`
	}{
		Error:    false,
		Message:  responseString,
		MatchIds: app.generateMatchIds(),
	}

	app.writeJSON(w, http.StatusAccepted, response)
}

func (app *Config) GetMatchDetailsBatch(w http.ResponseWriter, r *http.Request) {
	devId := chi.URLParam(r, "developerId")
	_ = chi.URLParam(r, "signature")
	_ = chi.URLParam(r, "session")
	timestamp := chi.URLParam(r, "timestamp")
	matchIds := strings.Split(chi.URLParam(r, "matchIds"), ",")
	matches := make([]Match, len(matchIds))
	for i := range matchIds {
		matches[i] = getMatchDetailsById(matchIds[i])
	}

	responseString := fmt.Sprintf("Dev %s requesting %d Match IDs @ %v", devId, len(matchIds), timestamp)

	payload := jsonResponse{
		Error:   false,
		Message: responseString,
		Data:    matches,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func getMatchDetailsById(matchId string) Match {
	return generateMatch(matchId)
}
