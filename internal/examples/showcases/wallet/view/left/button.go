package left

import (
	"github.com/dsandor/qt/quick"

	_ "github.com/dsandor/qt/internal/examples/showcases/wallet/view/left/controller"
)

func init() { buttonTemplate_QmlRegisterType2("LeftTemplate", 1, 0, "ButtonTemplate") }

type buttonTemplate struct {
	quick.QQuickItem

	_ func(string) `signal:"clicked,->(controller.NewButtonController(nil))"`
}
