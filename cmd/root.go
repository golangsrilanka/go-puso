package cmd

import (
	"github.com/GolangSriLanka/go-puso/cmd/view"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/GolangSriLanka/go-puso/cmd/server"
)

var rootCmd = &cobra.Command{
	Use:   "go-puso",
	Short: "template repo",
	Long:  `Golang template repo form Golang Sri Lanka `,
}

func init() {
	rootCmd.AddCommand(server.RunServerCmd)
	rootCmd.AddCommand(view.ViewCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
