package widgets

import (
	termuiw "github.com/gizak/termui/v3/widgets"
	"github.com/shirou/gopsutil/v3/host"
	"log"
	"os"
)

type HostWidget struct {
	*termuiw.Paragraph
}

func (w *HostWidget) init() {
	w.Title = "Host"
	users, err := host.Users()
	if err != nil {
		log.Fatal(err)
	}
	user := users[0]
	hostname, _ := os.Hostname()
	w.Text = user.User + "@" + hostname
}

func (w *HostWidget) update() {
}

func NewHostWidget() *HostWidget {
	self := HostWidget{Paragraph: termuiw.NewParagraph()}
	self.init()
	return &self
}
