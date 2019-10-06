package main

import (
	"fmt"
	"github.com/inkyblackness/imgui-go"
	"github.com/skratchdot/open-golang/open"
	"github.com/tobyxdd/go-ping/monitor"
	"math"
	"os"
)

const mainWindowFlags = imgui.WindowFlagsNoTitleBar | imgui.WindowFlagsNoCollapse | imgui.WindowFlagsNoMove |
	imgui.WindowFlagsNoResize | imgui.WindowFlagsMenuBar | imgui.WindowFlagsNoBringToFrontOnFocus

var mainWindowStates struct {
	AddText string
}

func mainWindow(appCtx *appContext) {
	imgui.SetNextWindowPos(imgui.Vec2{0, 0})
	imgui.SetNextWindowSize(imgui.Vec2{X: appCtx.Platform.DisplaySize()[0], Y: appCtx.Platform.DisplaySize()[1]})
	currentTheme.PushMainWindowStyle()

	imgui.BeginV("Main", nil, mainWindowFlags)

	// Main Menu
	if imgui.BeginMenuBar() {
		if imgui.BeginMenu("File") {
			if imgui.MenuItem("Quit") {
				os.Exit(0)
			}
			imgui.EndMenu()
		}
		if imgui.BeginMenu("Options") {
			if imgui.MenuItem("Edit config (Needs restarting)") {
				_ = saveConfig(appCtx.AppConfig)
				_ = open.Start(appConfigFilename)
				os.Exit(0)
			}
			imgui.EndMenu()
		}
		imgui.EndMenuBar()
	}

	metricsMap := appCtx.Monitor.Export()

	for _, t := range appCtx.AppConfig.Targets {
		drawTargetItem(appCtx, &t, metricsMap[t.Name])
	}

	imgui.PushItemWidth(imgui.ContentRegionAvail().X)
	imgui.InputText("", &mainWindowStates.AddText)
	imgui.PopItemWidth()
	if imgui.ButtonV("Add", imgui.Vec2{X: imgui.ContentRegionAvail().X}) {
		addTarget(appCtx, mainWindowStates.AddText)
	}

	imgui.End()

	currentTheme.PopMainWindowStyle()
}

func padLTS(lts []float32, historySize int) []float32 {
	if len(lts) >= historySize {
		return lts
	} else {
		pad := make([]float32, historySize-len(lts))
		return append(pad, lts...)
	}
}

func drawTargetItem(appCtx *appContext, t *target, metrics *monitor.Metrics) {
	imgui.Text(t.Name)
	if imgui.BeginPopupContextItemV(t.Name, 1) {
		var switchText string
		if t.Enabled {
			switchText = "Disable"
		} else {
			switchText = "Enable"
		}
		if imgui.MenuItem(switchText) {
			setTargetEnableState(appCtx, t.Name, !t.Enabled)
		}
		if imgui.MenuItem("Remove") {
			removeTarget(appCtx, t.Name)
			imgui.EndPopup()
			return
		}
		imgui.EndPopup()
	}

	if !t.Enabled {
		imgui.SameLine()
		imgui.Text("(Disabled)")
		imgui.Separator()
		return
	} else {
		if addr := appCtx.InfoMap[t.Name].Addr; addr != nil {
			addrString := addr.String()

			// Append to the previous menu here by using the same name
			if imgui.BeginPopupContextItemV(t.Name, 1) {
				imgui.Separator()
				if imgui.MenuItem("Lookup " + addrString) {
					_ = open.Start(fmt.Sprintf(appCtx.AppConfig.UI.IPLookupURL, addrString))
				}
				imgui.EndPopup()
			}

			if addr.String() != t.Name {
				imgui.SameLine()
				imgui.Text(fmt.Sprintf("(%s)", addrString))
			}
		}
	}

	if metrics == nil {
		// Enabled but no metrics, error?
		if err := appCtx.InfoMap[t.Name].Error; err != nil {
			imgui.PushStyleColor(imgui.StyleColorText, currentTheme.LatencyColors()[2])
			imgui.Text(err.Error())
			imgui.PopStyleColor()
		}
		imgui.Separator()
		return
	}
	lts := make([]float32, len(metrics.Results))
	for i := range metrics.Results {
		lts[i] = float32(metrics.Results[i].RTT.Nanoseconds()) / 1e6
	}
	colorIndex := 0
	if metrics.Median > float32(appCtx.AppConfig.Ping.LatencyYellowThreshold) {
		if metrics.Median > float32(appCtx.AppConfig.Ping.LatencyRedThreshold) {
			colorIndex = 2
		} else {
			colorIndex = 1
		}
	}
	imgui.PushStyleColor(imgui.StyleColorText, currentTheme.LatencyColors()[colorIndex])
	imgui.Text(fmt.Sprintf("Median: %.2fms, Mean: %.2fms, Loss: %.2f%%", metrics.Median, metrics.Mean, float32(metrics.PacketsLost)/float32(metrics.PacketsSent)*100))
	imgui.Text(fmt.Sprintf("Best: %.2fms, Worst: %.2fms, StdDev: %.2fms", metrics.Best, metrics.Worst, metrics.StdDev))
	imgui.PopStyleColor()

	imgui.PlotHistogramV("", padLTS(lts, appCtx.AppConfig.Ping.HistorySize), 0,
		"", 0, math.MaxFloat32, imgui.Vec2{imgui.ContentRegionAvail().X, 100})
	imgui.Separator()
}
