package chat

import (
	"fmt"
	"zodream/modules/chat/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

func Register(app iris.Party) {
	app.Get("/", controllers.Index)
	setupWebsocket(app)
}

func setupWebsocket(app iris.Party) {
	// create our echo websocket server
	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})
	ws.OnConnection(handleConnection)

	// register the server on an endpoint.
	// see the inline javascript code in the websockets.html,
	// this endpoint is used to connect to the server.
	app.Get("/ws", ws.Handler())
	// serve the javascript built'n client-side library,
	// see websockets.html script tags, this path is used.
	app.Any("/iris-ws.js", websocket.ClientHandler())
}

func handleConnection(c websocket.Connection) {
	// Read events from browser
	c.On("chat", func(msg string) {
		// Print the message to the console, c.Context() is the iris's http context.
		fmt.Printf("%s sent: %s\n", c.Context().RemoteAddr(), msg)
		// Write message back to the client message owner with:
		// c.Emit("chat", msg)
		// Write message to all except this client with:
		c.To(websocket.Broadcast).Emit("chat", msg)
	})
}
