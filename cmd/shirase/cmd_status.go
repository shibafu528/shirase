package main

import (
	"github.com/shibafu528/shirase/cmd/shirase/status"
)

func init() {
	rootCmd.AddCommand(status.Cmd)
}
