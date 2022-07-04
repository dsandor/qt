package d

import (
	"github.com/dsandor/qt/core"

	_ "github.com/dsandor/qt/internal/cmd/moc/test/sub/conf"
)

type StructSubGoD struct{}
type StructSubMocD struct{ core.QObject }
