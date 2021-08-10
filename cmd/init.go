package cmd

import (
	"log"

	dr "github.com/aedwa038/driveBackup/drive"
	"github.com/spf13/cobra"
)

var cmdInit = &cobra.Command{
	Use:   "init",
	Short: "init backup utility",
	Long:  ``,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		initDrive(cmd)
	},
}

func init() {
	rootCmd.AddCommand(cmdInit)
}

func initDrive(cmd *cobra.Command) {
	ctx := cmd.Context()
	_, err := dr.NewClient(ctx, "credentials.json")
	if err != nil {
		log.Fatal("Error initializing app %v", err)
	}

}
