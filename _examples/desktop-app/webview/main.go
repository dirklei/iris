package main

import (
	"github.com/kataras/iris/v12"
	"github.com/zserge/webview"
)

const addr = "127.0.0.1:8080"

/*
	# Windows requires special linker flags for GUI apps.
	# It's also recommended to use TDM-GCC-64 compiler for CGo.
	# http://tdm-gcc.tdragon.net/download
	#
	#
	$ go build -ldflags="-H windowsgui" -o myapp.exe # build for windows
	$ ./myapp.exe # run
	#
	#
	# Note: if you see "use option -std=c99 or -std=gnu99 to compile your code"
	# please refer to: https://github.com/zserge/webview/issues/188
*/
func main() {
	go runServer()
	showAndWaitWindow()
}

func runServer() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1> Hello Desktop</h1>")
	})
	app.Listen(addr)
}

func showAndWaitWindow() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("https://iris-go.com")
	w.Run()
}
