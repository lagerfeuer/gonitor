package widgets

import (
	"fmt"
	termuiw "github.com/gizak/termui/v3/widgets"
	"github.com/shirou/gopsutil/v3/mem"
	"gonitor/util"
	"log"
	"time"
)

type MemoryWidget struct {
	*termuiw.Gauge
	interval time.Duration
}

func (w *MemoryWidget) init() {
	w.Title = "Memory"
	w.TitleStyle = util.ClearBold

	if w.interval > 0 {
		go func() {
			for range time.NewTicker(w.interval).C {
				w.update()
			}
		}()
	}
}

func (w *MemoryWidget) update() {
	memory, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalf("Error reading memory data: %v", err)
	}

	w.Percent = int(memory.UsedPercent)
	gb := uint64(1024 * 1024 * 1024)
	w.Title = fmt.Sprintf("Memory (%dG/%dG used)", memory.Used/gb, memory.Total/gb)
}

func NewMemoryWidget(interval time.Duration) *MemoryWidget {
	widget := MemoryWidget{
		Gauge:    termuiw.NewGauge(),
		interval: interval,
	}
	widget.init()
	return &widget
}
