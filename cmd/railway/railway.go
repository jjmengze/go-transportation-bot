package main

import (
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"log"
	"math/rand"
	"os"
	"time"
)

// KlogWriter serves as a bridge between the standard log package and the glog package.
type KlogWriter struct{}

// Write implements the io.Writer interface.
func (writer KlogWriter) Write(data []byte) (n int, err error) {
	klog.InfoDepth(1, string(data))
	return len(data), nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	defer klog.Flush()
	log.SetOutput(KlogWriter{})
	log.SetFlags(0)
	command := NewAPIServerCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

func NewAPIServerCommand() *cobra.Command {
	serverFlags := NewServerRunOptions()
	cmd := &cobra.Command{
		Use:          "railway-api-server",
		Long:         `The railway API server provide Taiwan railway time query function.`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	serverFlags.AddFlags(cmd.Flags())
	return cmd
}
