package main

import (
	"os"

	"github.com/dsandor/qt/widgets"

	"github.com/dsandor/qt/internal/examples/uitools/calculator/ui"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	ui.NewCalculatorForm(nil).Show()

	widgets.QApplication_Exec()
}
