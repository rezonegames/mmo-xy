package main

import (
	"math/rand"
	"mmo-xy/config"
	"mmo-xy/internal/game"
	"mmo-xy/internal/web"
	"mmo-xy/models"
	"mmo-xy/pkg/log"
	"os"
	"sync"

	"github.com/urfave/cli"
	"golang.org/x/net/context"
)

func main() {
	app := cli.NewApp()
	app.Name = "xy"
	app.Author = "rezone games"
	app.Version = "0.0.1"
	app.Usage = "xy"
	app.Action = serve
	app.Run(os.Args)
}

func serve(ctx *cli.Context) error {
	rand.New(rand.NewSource(99))
	ctx1 := context.Background()
	log.InitLog()
	config.InitConfig()
	sc := config.ServerConfig
	models.StartMongo(ctx1, sc.Mongo.Uri, sc.Mongo.Db, false)
	models.StartRedis(ctx1, sc.Redis.Host, sc.Redis.Db, "")
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() { defer wg.Done(); web.StartUp() }()  // 开启web服务器
	go func() { defer wg.Done(); game.StartUp() }() // 开启游戏服
	wg.Wait()
	return nil
}
