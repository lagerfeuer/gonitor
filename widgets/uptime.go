package widgets

import (
	"fmt"
	termuiw "github.com/gizak/termui/v3/widgets"
	"github.com/shirou/gopsutil/v3/host"
	"gonitor/util"
	"time"
)

type UptimeWidget struct {
	*termuiw.Paragraph
	interval time.Duration
}

func (w *UptimeWidget) init() {
	w.Title = "Uptime"
	w.update()

	if w.interval > 0 {
		go func() {
			for range time.NewTicker(w.interval).C {
				w.update()
			}
		}()
	}
}
func (w *UptimeWidget) update() {
	uptime, err := host.Uptime()
	if err != nil {
		w.Text = fmt.Sprintf("Error fetching uptime: %v", err)
	} else {
		w.Text = util.UptimeToHumanReadable(uptime)
	}
}

func NewUptimeWidget(interval time.Duration) *UptimeWidget {
	self := UptimeWidget{
		Paragraph: termuiw.NewParagraph(),
		interval: interval,
	}
	self.init()
	return &self
}
