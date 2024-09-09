/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/mselser95/microservice-template/api"
	"github.com/mselser95/microservice-template/internal/config"
	"github.com/mselser95/microservice-template/internal/server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the service",
	Long:  `Starts the service`,
	Run: func(cmd *cobra.Command, args []string) {

		// Get Environment name
		env := os.Getenv("ENVIRONMENT")
		if env == "" {
			env = "dev"
		}

		logger, _ := zap.NewProduction()

		var path string
		switch env {
		case "dev":
			path = "./app.config.json"
		default:
			path = "/etc/app/app.config.json"
		}

		// Load the configuration
		cfg, err := config.LoadConfig(path)
		if err != nil {
			logger.Error("error-getting-config", zap.Error(err))
			panic(err)
		}
		_ = cfg

		// Now we create the GQL client

		// Now we load the GRPC server
		s := server.NewServer()
		go server.StartServer(*cfg, s, logger)

		// Now we load the API server
		go api.NewApiServer(*cfg, logger)

		// Wait forever
		select {}

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
