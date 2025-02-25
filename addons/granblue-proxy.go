package addons

import (
	"bytes"
	"log"
	"strings"
	"time"

	"mitoma-proxy-go/constants"
	"mitoma-proxy-go/utils"

	"github.com/lqqyt2423/go-mitmproxy/proxy"
)

const (
	RESULT_JSON  = "result.json"
	CREATE_QUEST = "create_quest"
	WIN_STR      = "\"win\""
)

func (a *GranblueAddon) RecordLapTime() {
	currentTime := time.Now()

	if a.lastWinTime != nil {
		elapsed := currentTime.Sub(*a.lastWinTime)
		log.Printf("ラップタイム: %.3f 秒", elapsed.Seconds())
	}

	a.lastWinTime = &currentTime
}

type GranblueAddon struct {
	proxy.BaseAddon
	lastWinTime *time.Time
}

func (a *GranblueAddon) Request(f *proxy.Flow) {
	if f.Request == nil {
		return
	}

	if strings.Contains(f.Request.URL.Path, CREATE_QUEST) {
		// SendKeyPress(constants.VK_F5)
		utils.SendKeyPress(constants.VK_F22)
	}
}

func (a *GranblueAddon) Response(f *proxy.Flow) {
	if f.Response == nil || f.Response.Body == nil {
		return
	}

	reqPath := f.Request.URL.Path
	responseBodyByte := f.Response.Body
	if strings.Contains(reqPath, RESULT_JSON) {
		responseBody := utils.ReadGzip(responseBodyByte)

		if bytes.Contains(responseBody, []byte(WIN_STR)) {
			utils.SendKeyPress(constants.VK_F13)
			a.RecordLapTime()
		}
		return
	}
}
