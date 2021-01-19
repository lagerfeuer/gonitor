package main

import (
	ui "github.com/gizak/termui/v3"
	custom "gonitor/widgets"
	"log"
	"os"
	"runtime"
	"time"
)

const (
	updateInterval = time.Second
	cpuColumns     = 2
)

type Bounds struct {
	start int
	end   int
}

type Dashboard struct {
	grid *ui.Grid
	bar  *ui.Grid

	host   *custom.HostWidget
	uptime *custom.UptimeWidget
	help   *custom.HelpWidget

	cpu      *custom.CPUWidget
	memory   *custom.MemoryWidget
	procList *custom.ProcessListWidget

	drawables []ui.Drawable
}

var dashboard *Dashboard

func initWidgets() {
	dashboard.host = custom.NewHostWidget()
	dashboard.uptime = custom.NewUptimeWidget(updateInterval)
	dashboard.help = custom.NewHelpWidget()

	dashboard.cpu = custom.NewCPUWidget(updateInterval, cpuColumns)
	dashboard.memory = custom.NewMemoryWidget(updateInterval)
	dashboard.procList = custom.NewProcessListWidget(updateInterval)

	log.Println("Widgets initialized")
}

func initUI() {
	//half := 1.0 / 2
	third := 1.0 / 3
	//fourth := 1.0 / 4

	dashboard.bar = ui.NewGrid()
	dashboard.bar.Set(
		ui.NewRow(1,
			ui.NewCol(third, dashboard.host),
			ui.NewCol(third, dashboard.uptime),
			ui.NewCol(third, dashboard.help),
		),
	)

	dashboard.grid = ui.NewGrid()

	width, height := ui.TerminalDimensions()
	barHeight := 3
	dashboard.bar.SetRect(0, 0, width, barHeight)
	cpuBounds := Bounds{barHeight, barHeight + runtime.NumCPU()/cpuColumns + 2}
	dashboard.cpu.SetRect(0, cpuBounds.start, width, cpuBounds.end)
	memBounds := Bounds{cpuBounds.end, cpuBounds.end + 3}
	dashboard.memory.SetRect(0, memBounds.start, width, memBounds.end)
	procBounds := Bounds{memBounds.end, height}
	dashboard.procList.SetRect(0, procBounds.start, width, procBounds.end)

	dashboard.drawables = []ui.Drawable{dashboard.bar, dashboard.cpu, dashboard.memory, dashboard.procList}
	log.Println("Dashboard initialized")
}

func loop() {
	ui.Render(dashboard.bar)
	defer ui.Clear()

	ticker := time.NewTicker(updateInterval).C
	uiEvents := ui.PollEvents()

	for {
		// TODO handle signals
		select {

		case <-ticker:
			for _, drawable := range dashboard.drawables {
				ui.Render(drawable)
			}

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
