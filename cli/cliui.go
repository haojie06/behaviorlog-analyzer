package cli

import (
	"behaviorlog-analyzer/utils"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// 如果暂时不使用termui不要再init里面进行初始化，会导致标准输入失效

// func init() {
// 	err := ui.Init()
// 	utils.CheckErr(err, "UI初始化")
// }

func Start() {
	err := ui.Init()
	utils.CheckErr(err, "UI初始化")
	p := widgets.NewParagraph()
	p.Text = "Behavior Log Analyzer!"
	p.SetRect(0, 0, 25, 5)
	ui.Render(p)
}
