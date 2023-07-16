package account

import (
	"fmt"

	"github.com/shibafu528/shirase"
	"github.com/shibafu528/shirase/db"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [USERNAME]",
	Short: "Create new account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, q := db.Open()
		key, err := shirase.NewKeyPair()
		if err != nil {
			panic(err)
		}
		res, err := q.CreateAccount(cmd.Context(), db.CreateAccountParams{
			Username:   args[0],
			PrivateKey: string(key.PrivateKey),
			PublicKey:  string(key.PublicKey),
		})
		if err != nil {
			panic(err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID=%d\n", id)
	},
}

func init() {
	Cmd.AddCommand(createCmd)
}
