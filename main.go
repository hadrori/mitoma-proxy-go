package main

import (
	"context"
	"fmt"
	"log"
	"syscall"

	"mitoma-proxy-go/addons"

	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

var (
	user32          = syscall.NewLazyDLL("user32.dll")
	procKeybd_Event = user32.NewProc("keybd_event")
)

const (
	PORT         = 8080
	RESULT_JSON  = "result.json"
	CREATE_QUEST = "create_quest"
	WIN_STR      = "\"win\""
)

func main() {
	ctx := context.Background()

	opts := &proxy.Options{
		Addr:              fmt.Sprintf(":%d", PORT),
		StreamLargeBodies: 1024 * 1024 * 5,
	}

	p, err := proxy.NewProxy(opts)
	if err != nil {
		log.Fatal(err)
	}

	p.AddAddon(&addons.GranblueAddon{})

	fmt.Printf("Starting mitmproxy-go on port %d...\n", PORT)
	fmt.Println("初回実行時は証明書の設定が必要です:")
	fmt.Printf("証明書ファイル %s をシステムにインポートしてください\n", "~/.mitmproxy/mitmproxy-ca-cert.pem")

	err = p.Start()
	if err != nil {
		log.Fatal(err)
	}

	<-ctx.Done()
}
