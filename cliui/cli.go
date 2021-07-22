package cliui

import (
	"behaviorlog-analyzer/utils"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func init() {
	err := ui.Init()
	utils.CheckErr(err, "UI初始化")
}

func Start() {
	p := widgets.NewParagraph()
	p.Text = "Behavior Log Analyzer!"
	p.SetRect(0, 0, 25, 5)
	ui.Render(p)
}

func PrintText()
