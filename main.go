package main

import (
	ui "github.com/gizak/termui/v3"
	custom "gonitor/widgets"
	"log"
	"os"
	"time"
)

const (
	updateInterval = time.Second
)

type Dashboard struct {
	grid *ui.Grid

	host   *custom.HostWidget
	uptime *custom.UptimeWidget
	help   *custom.HelpWidget
}

var dashboard *Dashboard

func initWidgets() {
	dashboard.host = custom.NewHostWidget()
	dashboard.uptime = custom.NewUptimeWidget(updateInterval)
	dashboard.help = custom.NewHelpWidget()

	log.Println("Widgets initialized")
}

func initUI() {
	third := 1.0 / 3

	dashboard.grid = ui.NewGrid()
	dashboard.grid.Set(
		ui.NewRow(1.0/5,
			ui.NewCol(third, dashboard.host),
			ui.NewCol(third, dashboard.uptime),
			ui.NewCol(third, dashboard.help),
		),
	)

	width, height := ui.TerminalDimensions()
	dashboard.grid.SetRect(0, 0, width, height)

	log.Println("Dashboard initialized")
}

func loop() {
	ui.Render(dashboard.grid)

	updateTimer := time.NewTicker(updateInterval).C
	uiEvents := ui.PollEvents()

	for {
		// TODO handle signals
		select {

		case <-updateTimer:
			ui.Render(dashboard.grid)

		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				os.Exit(0)
			}
		}
	}
}

func main() {
	// TODO: parse args

	file, err := os.OpenFile("/tmp/gonitor.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	log.SetOutput(file)
	log.Println("Logging initialized")
	dashboard = &Dashboard{}

	initWidgets()
	initUI()

	loop()
}
