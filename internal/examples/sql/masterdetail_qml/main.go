//go:build !qml
// +build !qml

package main

import (
	"os"

	"github.com/dsandor/qt/core"
	"github.com/dsandor/qt/widgets"

	"github.com/dsandor/qt/internal/examples/sql/masterdetail_qml/controller"

	"github.com/dsandor/qt/internal/examples/sql/masterdetail_qml/view"
)

func main() {
	qApp := widgets.NewQApplication(len(os.Args), os.Args)

	controller.NewController(nil).InitWith(core.NewQFile2(":/albumdetails.xml"), qApp)

	view.NewViewController(nil, 0).Show()

	qApp.Exec()
}
