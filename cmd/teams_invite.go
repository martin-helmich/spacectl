package cmd

import (
	"fmt"

	"errors"
	"github.com/gosuri/uitable"
	"github.com/mittwald/spacectl/client/teams"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var teamInviteFlags struct {
	Email   string
	UserID  string
	Message string
	Role    string
}

// inviteCmd represents the invite command
var teamInviteCmd = &cobra.Command{
	Use:   "invite -t <team-id> -e <email> -m <message>",
	Short: "Invite new users to your team",
	Long:  `Invite a new user into your team`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var invite teams.Invite

		teamID := viper.GetString("teamID")

		if teamID == "" {
			return errors.New("must provide team (--team-id or -t)")
		}

		if teamInviteFlags.Email == "" && teamInviteFlags.UserID == "" {
			return errors.New("must provide user (either --email|-e or --user-id|-u)")
		}

		userTemplate := "inviting user \"%s\" into team %s\n"
		if teamInviteFlags.UserID != "" {
			fmt.Printf(userTemplate, teamInviteFlags.UserID, teamID)
			invite, err = api.Teams().InviteByUID(
				teamID,
				teamInviteFlags.UserID,
				teamInviteFlags.Message,
				teamInviteFlags.Role,
			)
		} else if teamInviteFlags.Email != "" {
			fmt.Printf(userTemplate, teamInviteFlags.Email, teamID)
			invite, err = api.Teams().InviteByEmail(
				teamID,
				teamInviteFlags.Email,
				teamInviteFlags.Message,
				teamInviteFlags.Role,
			)
		}

		if err != nil {
			return err
		}

		fmt.Printf("invite %s issued\n", invite.ID)

		table := uitable.New()
		table.MaxColWidth = 80
		table.Wrap = true

		table.AddRow("ID:", invite.ID)
		table.AddRow("Message:", invite.Message)
		table.AddRow("State:", invite.State)

		fmt.Println(table)

		return nil
	},
}

func init() {
	teamsCmd.AddCommand(teamInviteCmd)

	teamInviteCmd.Flags().StringVarP(&teamInviteFlags.Email, "email", "e", "", "Email address of the user to invite")
	teamInviteCmd.Flags().StringVarP(&teamInviteFlags.UserID, "user-id", "u", "", "User ID of the user to invite")
	teamInviteCmd.Flags().StringVarP(&teamInviteFlags.Message, "message", "m", "", "Invitation message")
	teamInviteCmd.Flags().StringVarP(&teamInviteFlags.Role, "role", "r", "", "User role")
}
