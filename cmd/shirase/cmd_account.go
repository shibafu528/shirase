package main

import (
	"github.com/shibafu528/shirase/cmd/shirase/account"
)

func init() {
	rootCmd.AddCommand(account.Cmd)
}
