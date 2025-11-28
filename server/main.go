package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cargoreligion/booking/server/api"
	"github.com/cargoreligion/booking/server/infrastructure/db"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Configure zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Info().Msg("Starting HTTP server...")

	log.Info().Msg("Trying to connect to database...")
	dbInst, err := db.GetDbConnection(5)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	defer dbInst.Close()

	dbc := db.NewDbClient(dbInst)

	router := api.NewRouter(dbc)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	log.Fatal().Err(http.ListenAndServe(port, router)).Msg("Server failed to start.")
}
