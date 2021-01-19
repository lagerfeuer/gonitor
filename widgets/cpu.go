package widgets

import (
	"fmt"
	termuiw "github.com/gizak/termui/v3/widgets"
	"github.com/shirou/gopsutil/v3/cpu"
	"gonitor/util"
	"log"
	"strings"
	"time"
)

type CPUWidget struct {
	*termuiw.Paragraph
	interval time.Duration
	count    int
	data     []float64

	last    []cpu.TimesStat
	current []cpu.TimesStat

	columns int
}

func (w *CPUWidget) init() {
	w.Title = "CPU"
	w.TitleStyle = util.ClearBold

	if w.interval > 0 {
		go func() {
			for range time.NewTicker(w.interval).C {
				w.update()
			}
		}()
	}
}

func (w *CPUWidget) update() {
	w.last = w.current
	w.current, _ = cpu.Times(true)
	for i := range w.data {
		total := w.current[i].Total() - w.last[i].Total()
		idle := w.current[i].Idle - w.last[i].Idle
		w.data[i] = (idle / total) * 100.0
	}

	width := w.Inner.Max.X - w.Inner.Min.X
	cols := w.columns
	rowwidth := ((width - 2) / cols) - len(" CPUXX[] ")
	log.Println("[CPU] width is", width)
	log.Println("[CPU] cols is", cols)
	log.Println("[CPU] rowwidth is", rowwidth)
	builder := strings.Builder{}
	for i := range w.data {
		load := 1 - (w.data[i] / 100)
		if i > 0 && i%cols == 0 {
			builder.WriteString("\n")
		}
		usedIndicatorLength := int(float64(rowwidth) * load)
		freeIndicatorLength := rowwidth - usedIndicatorLength
		builder.WriteString(fmt.Sprintf(" CPU%2d[%s%s] ",
			i+1,
			strings.Repeat("|", usedIndicatorLength),
			strings.Repeat(" ", freeIndicatorLength)))
	}
	w.Text = builder.String()
}

func NewCPUWidget(interval time.Duration, columns int) *CPUWidget {
	cpuTimes, err := cpu.Times(true)
	if err != nil {
		log.Fatalf("Error reading CPU times: %v", err)
	}

	widget := CPUWidget{
		Paragraph: termuiw.NewParagraph(),
		interval:  interval,
		count:     len(cpuTimes),
		data:      make([]float64, len(cpuTimes)),
		last:      cpuTimes,
		current:   cpuTimes,
		columns:   columns,
	}
	widget.init()
	return &widget
}
