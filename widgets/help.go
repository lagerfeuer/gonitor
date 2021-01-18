package widgets

import (
	termuiw "github.com/gizak/termui/v3/widgets"
)

type HelpWidget struct {
	*termuiw.Paragraph
}

func (w *HelpWidget) init() {
	w.Title = "Help"
	w.Text = "Press [q] to exit."
}

func (w *HelpWidget) update() {
}

func NewHelpWidget() *HelpWidget {
	self := HelpWidget{Paragraph: termuiw.NewParagraph()}
	self.init()
	return &self
}
