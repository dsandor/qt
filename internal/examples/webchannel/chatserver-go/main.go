package main

import (
	"os"

	"github.com/dsandor/qt/core"
	"github.com/dsandor/qt/network"
	"github.com/dsandor/qt/webchannel"
	"github.com/dsandor/qt/websockets"

	"github.com/dsandor/qt/internal/examples/webchannel/shared"
)

func main() {
	core.NewQCoreApplication(len(os.Args), os.Args)

	server := websockets.NewQWebSocketServer2("QWebChannel Standalone Example Server", websockets.QWebSocketServer__NonSecureMode, nil)
	if !server.Listen(network.NewQHostAddress9(network.QHostAddress__LocalHost), 12345) {
		panic("Failed to open web socket server")
	}

	clientWrapper := shared.NewWebSocketClientWrapper(nil)
	clientWrapper.Init(server)

	channel := webchannel.NewQWebChannel(nil)

	clientWrapper.ConnectClientConnected(func(client *shared.WebSocketTransport) {
		channel.ConnectTo(client.QWebChannelAbstractTransport_PTR())
	})

	chatserver := NewChatServer(nil)
	channel.RegisterObject("chatserver", chatserver)

	core.QCoreApplication_Exec()
}
