package main

import (
	"flag"
	"git.leantar.de/torge/fimserver/modules/config"
	"git.leantar.de/torge/fimserver/modules/preparation"
	"git.leantar.de/torge/fimserver/repository"
	"git.leantar.de/torge/fimserver/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	Server     server.Config     `yaml:"server"`
	Repository repository.Config `yaml:"repository"`
}

var (
	configPath = flag.String("config", "config.yaml", "Specify a path to load the config from")
	setupMode  = flag.Bool("setup", false, "Prepare the database of the application")
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Parse command line arguments
	flag.Parse()

	// Read config file
	var conf Config
	err := config.FromYamlFile(*configPath, &conf)
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("failed to read config")
	}

	if *setupMode {
		err := setup(conf)
		if err != nil {
			log.Fatal().Caller().Err(err).Msg("failed to run preparation")
		}
	} else {
		err := run(conf)
		if err != nil {
			log.Fatal().Caller().Err(err).Msg("failed to run server")
		}
	}
}

// Start the normal execution mode
func run(conf Config) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	repo := repository.New(conf.Repository)
	srv := server.New(repo, conf.Server)

	go func() {
		if err := srv.Run(); err != nil {
			log.Fatal().Caller().Err(err).Msg("server failed to run")
		}
	}()

	<-quit
	srv.Stop()
	return nil
}

// Run the setup mode. This creates all required casbin rules, an admin user and all relations inside the database.
func setup(conf Config) error {
	repo := repository.New(conf.Repository)

	return preparation.Setup(repo)
}
