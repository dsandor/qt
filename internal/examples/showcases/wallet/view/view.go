package view

import (
	"github.com/dsandor/qt/quick"

	_ "github.com/dsandor/qt/internal/examples/showcases/wallet/view/controller"
)

func init() { viewTemplate_QmlRegisterType2("ViewTemplate", 1, 0, "ViewTemplate") }

type viewTemplate struct {
	quick.QQuickItem

	_ func(b bool) `signal:"Blur,<-(controller.NewViewController(nil))"`
}
