package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/shibafu528/shirase"
	"github.com/shibafu528/shirase/router"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		// construct http server
		s := &http.Server{
			Addr:    shirase.GlobalConfig.HttpListenAddr(),
			Handler: router.New(),
		}

		// handle SIGINT, SIGTERM
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			sig := <-sigch
			log.Printf("received signal %v, exiting gracefully...", sig)
			if err := s.Shutdown(context.Background()); err != nil {
				log.Printf("error in shutdown server: %v", err)
			}
			wg.Done()
		}()

		// start http server
		log.Printf("http server started on %s", shirase.GlobalConfig.HttpListenAddr())
		if err := s.ListenAndServe(); err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
