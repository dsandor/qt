package swift

import "github.com/dsandor/qt/interop"

func init() {
	interop.ReturnPointersAsStrings = true
	interop.SupportsSyncCallsIntoRemote = false
}
