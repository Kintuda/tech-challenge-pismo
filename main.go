package main

import (
	"os"

	"github.com/kintuda/tech-challenge-pismo/cmd"
	"github.com/rs/zerolog/log"
)

func main() {
	root := cmd.NewRootCmd()

	if err := root.Execute(); err != nil {
		log.Error().Err(err).Msg("command resulted in error")
		os.Exit(1)
	}
}
