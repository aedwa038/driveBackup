package cmd

import (
	"fmt"
	"log"

	dr "github.com/aedwa038/driveBackup/drive"
	"github.com/spf13/cobra"
)

var searchInit = &cobra.Command{
	Use:   "search",
	Short: "init backup utility",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		search(cmd, args[0])
	},
}

func init() {
	rootCmd.AddCommand(searchInit)
}

func search(cmd *cobra.Command, term string) {

	ctx := cmd.Context()
	driveService, err := dr.NewClient(ctx, "credentials.json")
	if err != nil {
		log.Fatalf("Error initializing app %v", err)
	}

	q := fmt.Sprintf("name = '%s'", term)
	fmt.Println(term)
	r, err := driveService.Client.Files.List().Q(q).Fields("nextPageToken, files(*)").Do()
	if err != nil {
		log.Fatalf("Error reading files %v", err)
	}

	fmt.Println(len(r.Files))
	for _, i := range r.Files {
		fmt.Printf("%s, %s, %s\n", i.Name, i.Id, i.MimeType)
	}
}
