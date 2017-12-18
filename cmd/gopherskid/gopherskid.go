package main

import (
	"net/http"
	"time"

	"bytes"
	"fmt"
	"github.com/frycm/gopher-scapes/internal/platform/hgt"
	"github.com/frycm/gopher-scapes/internal/platform/render"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os"
)

var log *zap.Logger

func main() {
	var err error
	log, err = zap.NewDevelopment()
	if err != nil {
		panic(errors.Wrap(err, "could not initialise logging"))
	}
	defer log.Sync()

	r := mux.NewRouter()

	// {region:(?:N|S)\\d{2}(?:W|E)\\d{2}}
	r.HandleFunc("/slopeSteepMap/fixedRegion/{region}", fixedRegion).
		Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(errors.Wrap(err, "could not start HTTP server"))
	}
}

func fixedRegion(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	region := vars["region"]
	latitude := region[0:3]
	longitude := region[3:6]

	log.Info(fmt.Sprintf("New request for fixedRegion: %s%s", latitude, longitude))

	var tile hgt.Tile
	err := loadHeightMap(&tile, latitude, longitude)
	var image bytes.Buffer
	render.ToPNG(&image)
	response.Write(image.Bytes())
}

func loadHeightMap(tile *hgt.Tile, latitude, longitude string) error {
	source, err := os.Open(fmt.Sprintf("test_data/raw_hgt/%s%s.hgt", latitude, longitude))
	if err != nil {
		return errors.Wrapf(err, "could open source hgt source file for %s%s", latitude, longitude)
	}
	return hgt.Load(tile, source)
}
