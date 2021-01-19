package widgets

import (
	termuiw "github.com/gizak/termui/v3/widgets"
	"gonitor/util"
)

type HelpWidget struct {
	*termuiw.Paragraph
}

func (w *HelpWidget) init() {
	w.Title = "Help"
	w.TitleStyle = util.ClearBold
	w.Text = "Press [q] to exit."
}

func (w *HelpWidget) update() {
}

func NewHelpWidget() *HelpWidget {
	widget := HelpWidget{Paragraph: termuiw.NewParagraph()}
	widget.init()
	return &widget
}
