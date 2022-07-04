package controller

import (
	"github.com/dsandor/qt/core"
	"github.com/dsandor/qt/gui"
)

type logoController struct {
	core.QObject

	_ func() `signal:"clicked,auto"`
}

func (c *logoController) clicked() {
	gui.QDesktopServices_OpenUrl(core.NewQUrl3("https://example.com/", 0))
}
