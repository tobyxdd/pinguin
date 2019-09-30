package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const appConfigFilename = "config.json"

type uiConfig struct {
	FontFile string  `json:"font_file"`
	FontSize float32 `json:"font_size"`
}

type pingConfig struct {
	HistorySize            int `json:"history_size"`
	IntervalMS             int `json:"interval_ms"`
	TimeoutMS              int `json:"timeout_ms"`
	LatencyYellowThreshold int `json:"latency_yellow_threshold"`
	LatencyRedThreshold    int `json:"latency_red_threshold"`
}

type target struct {
	Enabled bool   `json:"enabled"`
	Name    string `json:"name"`
}

type appConfig struct {
	UI      uiConfig   `json:"ui"`
	Ping    pingConfig `json:"ping"`
	Targets []target   `json:"targets"`
}

var defaultAppConfig = appConfig{
	UI: uiConfig{
		FontFile: "font.ttf",
		FontSize: 15,
	},
	Ping: pingConfig{
		HistorySize:            60,
		IntervalMS:             1000,
		TimeoutMS:              1000,
		LatencyYellowThreshold: 100,
		LatencyRedThreshold:    200,
	},
	Targets: []target{
		{
			Enabled: true,
			Name:    "www.google.com",
		},
		{
			Enabled: false,
			Name:    "1.1.1.1",
		},
	},
}

func loadAppConfig() appConfig {
	fbs, err := ioutil.ReadFile(appConfigFilename)
	if err != nil {
		out, _ := json.Marshal(defaultAppConfig)
		_ = ioutil.WriteFile(appConfigFilename, out, os.FileMode(0666))
		return defaultAppConfig
	}
	var c appConfig
	err = json.Unmarshal(fbs, &c)
	if err != nil {
		return defaultAppConfig
	}
	return c
}
