package b

import (
	"github.com/dsandor/qt/core"

	_ "github.com/dsandor/qt/internal/cmd/moc/test/sub/conf"
)

type StructSubGoB struct{}
type StructSubMocB struct{ core.QObject }
