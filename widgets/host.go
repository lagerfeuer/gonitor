package widgets

import (
	termuiw "github.com/gizak/termui/v3/widgets"
	"github.com/shirou/gopsutil/v3/host"
	"gonitor/util"
	"log"
	"os"
)

type HostWidget struct {
	*termuiw.Paragraph
}

func (w *HostWidget) init() {
	w.Title = "Host"
	w.TitleStyle = util.ClearBold
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
	widget := HostWidget{Paragraph: termuiw.NewParagraph()}
	widget.init()
	return &widget
}
