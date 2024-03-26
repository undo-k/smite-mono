package main

import (
	"errors"

	"github.com/undo-k/smite-mono/protos/protos"
)

func (app *Config) retrieveFromCache(godName string) (*protos.God, error) {
	god, ok := app.GodCache[godName]
	if ok {
		return god, nil
	}

	return nil, errors.New("failed to retrieve from cache: god does not exist")
}

func (app *Config) cacheInsert(god *protos.God) {
	app.GodCache[god.Name] = god
}

func createGodCache() (map[string]*protos.God, error) {

	gc := make(map[string]*protos.God)

	return gc, nil
}
