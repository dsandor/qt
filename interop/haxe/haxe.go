package haxe

import "github.com/dsandor/qt/interop"

func init() {
	interop.ReturnPointersAsStrings = true
	interop.SupportsSyncCallsIntoRemote = false
}
