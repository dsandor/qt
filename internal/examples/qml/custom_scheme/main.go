package main

import (
	"os"

	"github.com/dsandor/qt/core"
	"github.com/dsandor/qt/network"
	"github.com/dsandor/qt/qml"
	"github.com/dsandor/qt/quick"
	"github.com/dsandor/qt/widgets"
)

func main() {

	widgets.NewQApplication(len(os.Args), os.Args)

	var view = quick.NewQQuickView(nil)

	var (
		factory = qml.NewQQmlNetworkAccessManagerFactory()
		manager *network.QNetworkAccessManager
	)

	factory.ConnectCreate(func(parent *core.QObject) *network.QNetworkAccessManager {
		if manager != nil {
			return manager
		}

		manager = network.NewQNetworkAccessManager(parent)

		manager.ConnectCreateRequest(func(op network.QNetworkAccessManager__Operation, req *network.QNetworkRequest, outgoingData *core.QIODevice) *network.QNetworkReply {
			if req.Url().Scheme() != "custom" {
				return manager.CreateRequestDefault(op, req, outgoingData)
			}

			return NewCustomReply(op, req)
		})

		return manager
	})

	view.Engine().SetNetworkAccessManagerFactory(factory)
	view.SetSource(core.NewQUrl3("qrc:///qml/myfile.qml", 0)) //can be custom:/// as well
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	widgets.QApplication_Exec()
}
