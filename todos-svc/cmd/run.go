package cmd

import (
	"fmt"
	"github.com/fester-io/todos-lib/log"
	"github.com/fester-io/todos-todos-svc/internal/kafka"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the service",
	Run: func(cmd *cobra.Command, args []string) {
		if err := run(cmd, args); err != nil {
			log.Fatal("failed to run the service", zap.Any("err", err))
		}
	},
}

func run(_ *cobra.Command, _ []string) error {
	if err := kafka.InitClient(); err != nil {
		return err
	}

	kafka.StartConsumer()

	done := make(chan struct{}, 1)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		fmt.Printf("signal caught: %s\n", <-sig)
		done <- struct{}{}
	}()

	// wait for signal
	<-done

	return nil
}

func init() {
	rootCmd.AddCommand(runCmd)
}
