package status

import (
	"fmt"
	"strconv"

	"github.com/shibafu528/shirase"
	"github.com/shibafu528/shirase/db"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [account_id] [text]",
	Short: "Create new status",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		aid, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic(err)
		}
		text := args[1]

		_, q := db.Open()
		account, err := q.GetAccount(cmd.Context(), aid)
		if err != nil {
			panic(err)
		}

		snowflakeID, err := shirase.GenerateSnowflakeID()
		if err != nil {
			panic(err)
		}

		res, err := q.CreateStatus(cmd.Context(), db.CreateStatusParams{
			ID:        snowflakeID.Int64(),
			AccountID: account.ID,
			Text:      text,
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
