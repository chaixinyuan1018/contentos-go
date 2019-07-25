package commands

import (
	"fmt"
	"github.com/coschain/cobra"
	ctrl "github.com/coschain/contentos-go/app"
	"github.com/coschain/contentos-go/cmd/multinodetester/commands/test"
	"github.com/coschain/contentos-go/common"
	"github.com/coschain/contentos-go/config"
	"github.com/coschain/contentos-go/consensus"
	"github.com/coschain/contentos-go/db/storage"
	"github.com/coschain/contentos-go/iservices"
	"github.com/coschain/contentos-go/mylog"
	"github.com/coschain/contentos-go/node"
	"github.com/coschain/contentos-go/p2p"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

var VERSION = "defaultVersion"

var TestCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test count",
		Short: "start cosd nodes",
		Run:   startNodes,
	}
	return cmd
}

func makeNode(name string) (*node.Node, node.Config) {
	var cfg node.Config
	cfg.Name = name
	confdir := filepath.Join(config.DefaultDataDir(), cfg.Name)
	viper.SetConfigFile(confdir + "/config.toml")
	err := viper.ReadInConfig()
	if err == nil {
		_ = viper.Unmarshal(&cfg)
	} else {
		fmt.Printf("fatal: not be initialized (do `init` first)\n")
		os.Exit(1)
	}

	if cfg.DataDir != "" {
		dir, err := filepath.Abs(cfg.DataDir)
		if err != nil {
			common.Fatalf("DataDir in cfg cannot be converted to absolute path")
		}
		cfg.DataDir = dir
	}
	cfg.P2P.RunningCodeVersion = VERSION
	app, err := node.New(&cfg)
	if err != nil {
		fmt.Println("Fatal: ", err)
		os.Exit(1)
	}
	return app, cfg
}

type emptyWriter struct{}

func (ew emptyWriter) Write(p []byte) (int, error) {
	return 0, nil
}
func startNodes(cmd *cobra.Command, args []string) {
	// _ is cfg as below process has't used
	_, _ = cmd, args
	cnt, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	nodes := make([]*node.Node, 0, cnt)
	sks := make([]string, 0, cnt)

	for i := 0; i < cnt; i++ {
		name := fmt.Sprintf("%s_%d", TesterClientIdentifier, i)
		app, cfg := makeNode(name)
		app.Log = mylog.Init(cfg.ResolvePath("logs"), cfg.LogLevel, 3600*24*7)
		app.Log.Out = &emptyWriter{}
		app.Log.Info("Cosd running version: ", VERSION)
		RegisterService(app, cfg)
		if err := app.Start(); err != nil {
			common.Fatalf("start node failed, err: %v\n", err)
		}
		nodes = append(nodes, app)
		sks = append(sks, cfg.Consensus.LocalBpPrivateKey)
	}

	stopCh := make(chan struct{})
	go func() {
		SIGSTOP := syscall.Signal(0x13) //for windows compile
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, SIGSTOP, syscall.SIGUSR1, syscall.SIGUSR2)
		for {
			s := <-sigc
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, SIGSTOP, syscall.SIGINT:
				for i := range nodes {
					nodes[i].Log.Infoln("Got interrupt, shutting down...")
				}
				close(stopCh)
				return
			default:
				return
			}
		}
	}()


	c, err := nodes[0].Service(iservices.ConsensusServerName)
	if err != nil {
		panic(err)
	}
	css := c.(iservices.IConsensus)

	time.Sleep(2 * time.Second)
	for i := 1; i < cnt; i++ {
		name := fmt.Sprintf("initminer%d", i)
		if err = test.CreateAcc(name, sks[i], sks[0], css); err != nil {
			panic(err)
		}
	}
	fmt.Printf("created %d accounts\n", cnt-1)

	time.Sleep(2 * time.Second)
	for i := 1; i < cnt; i++ {
		name := fmt.Sprintf("initminer%d", i)
		if err = test.RegesiterBP(name, sks[i], css); err != nil {
			panic(err)
		}
	}
	fmt.Printf("registered %d bp\n", cnt-1)

	monitor := test.NewMonitor(nodes)
	go monitor.Run()

	<-stopCh
	for i := range nodes {
		nodes[i].Log.Info("start complete")
		//nodes[i].Wait()
		nodes[i].Stop()
		nodes[i].Log.Info("app exit success")
	}

}

func RegisterService(app *node.Node, cfg node.Config) {
	_ = app.Register(iservices.DbServerName, func(ctx *node.ServiceContext) (node.Service, error) {
		return storage.NewGuardedDatabaseService(ctx, "./db/")
	})

	_ = app.Register(iservices.TxPoolServerName, func(ctx *node.ServiceContext) (node.Service, error) {
		return ctrl.NewController(ctx, app.Log)
	})

	_ = app.Register(iservices.ConsensusServerName, func(ctx *node.ServiceContext) (node.Service, error) {
		var s node.Service
		switch ctx.Config().Consensus.Type {
		case "DPoS":
			s = consensus.NewDPoS(ctx, app.Log)
		case "SABFT":
			s = consensus.NewSABFT(ctx, app.Log)
		default:
			s = consensus.NewDPoS(ctx, app.Log)
		}
		return s, nil
	})

	_ = app.Register(iservices.P2PServerName, func(ctx *node.ServiceContext) (node.Service, error) {
		return p2p.NewServer(ctx, app.Log)
	})
}