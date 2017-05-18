package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/gosuri/uitable"
)

var teamsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List teams",
	Long:  `Lists all teams that you have access to.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		teams, err := spaces.Teams().List()
		if err != nil {
			return err
		}

		table := uitable.New()
		table.MaxColWidth = 50
		table.AddRow("ID", "DNS NAME", "NAME")

		for _, team := range teams {
			table.AddRow(team.ID, team.DNSName, team.Name)
		}

		fmt.Println(table)

		return nil
	},
}

func init() {
	teamsCmd.AddCommand(teamsListCmd)
}
