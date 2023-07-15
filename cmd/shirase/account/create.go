package account

import (
	"fmt"

	"github.com/shibafu528/shirase"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [USERNAME]",
	Short: "Create new account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, q := shirase.GlobalConfig.DB()
		res, err := q.CreateAccount(cmd.Context(), args[0])
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
