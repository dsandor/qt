package dialog

import (
	"github.com/dsandor/qt/core"

	_ "github.com/dsandor/qt/internal/examples/showcases/wallet/wallet/dialog/controller"
)

func init() { recoverTemplate_QmlRegisterType2("DialogTemplate", 1, 0, "RecoverTemplate") }

type recoverTemplate struct {
	dialogTemplate

	_ func(string) *core.QVariant `slot:"recover,->(controller.Controller)"`
}
