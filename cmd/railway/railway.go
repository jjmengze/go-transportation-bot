package main

import (
	"github.com/spf13/cobra"
	"go-transportation-bot/cmd/railway/app"
	"go-transportation-bot/pkg/controller/railway"
	"go-transportation-bot/pkg/modules/cache"
	"go-transportation-bot/pkg/repository"
	"go-transportation-bot/pkg/service"
	"go-transportation-bot/pkg/utils/signal"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
	"log"
	"math/rand"
	"net"
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
	serverFlags := app.NewServerRunOptions()
	cmd := &cobra.Command{
		Use:          "railway-api-server",
		Long:         `The railway API server provide Taiwan railway time query function.`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			klog.Infof("Default serving IP: %v", serverFlags.IP)
			klog.Infof("Default serving port: %v", serverFlags.Port)
			klog.Infof("Default connection redis: %v", serverFlags.Redis)

			return Run(serverFlags, signal.SetupSignalContext().Done())
		},
	}

	serverFlags.AddFlags(cmd.Flags())
	return cmd
}

func Run(option *app.ServerRunOptions, stopCh <-chan struct{}) error {

	cacheManger := cache.GetManager()
	railwayRepo := repository.NewRailwayRepository(cacheManger.GetRedisClient("127.0.0.1:6379"))
	railwaySvc := service.NewRailwayService(railwayRepo)

	lis, err := net.Listen("tcp", option.IP+":"+option.Port)
	if err != nil {
		klog.ErrorS(err, "Grpc start error occur: %s")
	}
	s := grpc.NewServer()
	railway.New(s, railwaySvc)

	return s.Serve(lis)
}
