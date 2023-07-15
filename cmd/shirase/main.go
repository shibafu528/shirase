package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/shibafu528/shirase"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "shirase",
}

func main() {
	// load config
	err := godotenv.Load()
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("[shirase.Config] Load from environment")
		} else {
			panic(err)
		}
	} else {
		log.Println("[shirase.Config] Load from .env file")
	}
	if err := envconfig.Process("", &shirase.GlobalConfig); err != nil {
		log.Fatalf("[shirase.Config] Load error!! %v\n", err)
	}

	// dispatch command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
