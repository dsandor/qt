package c

import (
	"github.com/dsandor/qt/core"

	_ "github.com/dsandor/qt/internal/cmd/moc/test/sub/conf"
)

type StructSubGoC struct{}
type StructSubMocC struct{ core.QObject }
