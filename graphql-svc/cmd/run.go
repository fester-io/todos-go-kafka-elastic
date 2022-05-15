package cmd

import (
	"fmt"
	"github.com/fester-io/todos-gql-kafka/internal/kafka"
	"github.com/fester-io/todos-gql-kafka/internal/server"
	"github.com/fester-io/todos-lib/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"strconv"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := run(cmd, args); err != nil {
			log.Fatal("failed to run", zap.Any("err", err))
		}
	},
}

func run(_ *cobra.Command, _ []string) error {
	if err := kafka.InitClient(); err != nil {
		return err
	}

	kafka.StartConsumer()

	fmt.Printf("starting server on port %s", strconv.FormatInt(8080, 10))
	if err := server.Start(8080); err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
