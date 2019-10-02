package main

import (
	"fmt"
	"github.com/inkyblackness/imgui-go"
	"github.com/tobyxdd/go-ping"
	"github.com/tobyxdd/go-ping/monitor"
	"github.com/tobyxdd/pinguin/platform"
	"github.com/tobyxdd/pinguin/renderer"
	"log"
	"net"
	"sync"
	"time"
)

const appVersion = "1.0.0"

var currentTheme theme = &lightTheme{}

type clipboard struct {
	platform platform.Platform
}

func (board clipboard) Text() (string, error) {
	return board.platform.ClipboardText()
}

func (board clipboard) SetText(text string) {
	board.platform.SetClipboardText(text)
}

func main() {
	// imgui init
	imCtx := imgui.CreateContext(nil)
	defer imCtx.Destroy()

	io := imgui.CurrentIO()

	// Platform
	plt, err := platform.NewGLFWPlatform(io, 300, 450, fmt.Sprintf("Pinguin %s", appVersion))
	if err != nil {
		log.Fatalln(err)
	}
	defer plt.Dispose()
	config := loadAppConfig()
	io.Fonts().AddFontFromFileTTF(config.UI.FontFile, config.UI.FontSize)
	io.SetClipboard(clipboard{platform: plt})
	io.SetIniFilename("")

	// Renderer
	rdr, err := renderer.NewOpenGL3Renderer(io)
	if err != nil {
		log.Fatalln(err)
	}
	defer rdr.Dispose()

	// Ping
	pgr, err := ping.New("0.0.0.0", "::")
	if err != nil {
		log.Fatalln(err)
	}
	mon := monitor.New(pgr, time.Duration(config.Ping.IntervalMS)*time.Millisecond, time.Duration(config.Ping.TimeoutMS)*time.Millisecond)
	mon.HistorySize = config.Ping.HistorySize

	appCtx := &appContext{
		AppConfig: config,
		Platform:  plt,
		Renderer:  rdr,
		Pinger:    pgr,
		Monitor:   mon,
		InfoMap:   make(map[string]targetInfo),
	}

	// Add targets from config
	for _, t := range config.Targets {
		if t.Enabled {
			ct := t // Copy the value to bypass go routine concurrency issue
			go func() {
				addr, err := net.ResolveIPAddr("ip", ct.Name)

				appCtx.InfoMapMutex.Lock()
				if err != nil {
					appCtx.InfoMap[ct.Name] = targetInfo{Error: err}
					return
				}
				appCtx.InfoMap[ct.Name] = targetInfo{Addr: addr}
				appCtx.InfoMapMutex.Unlock()

				_ = mon.AddTarget(ct.Name, *addr)
			}()
		}
	}

	// Run loop
	run(appCtx)
}

type appContext struct {
	AppConfig appConfig
	Platform  platform.Platform
	Renderer  renderer.Renderer

	// Ping stuff
	Pinger       *ping.Pinger
	Monitor      *monitor.Monitor
	InfoMap      map[string]targetInfo
	InfoMapMutex sync.RWMutex
}

type targetInfo struct {
	Error error
	Addr  *net.IPAddr
}

func run(appCtx *appContext) {
	currentTheme.SetStyleColors(imgui.CurrentStyle())

	for !appCtx.Platform.ShouldStop() {
		appCtx.Platform.ProcessEvents()
		appCtx.Platform.NewFrame()
		imgui.NewFrame()

		mainWindow(appCtx)

		imgui.Render()
		appCtx.Renderer.PreRender([4]float32{})
		appCtx.Renderer.Render(appCtx.Platform.DisplaySize(), appCtx.Platform.FramebufferSize(), imgui.RenderedDrawData())
		appCtx.Platform.PostRender()
	}

	_ = saveConfig(appCtx.AppConfig)
}
