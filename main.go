package main

import (
	"changeme/app/service"
	"embed"
	_ "embed"
	"github.com/wailsapp/wails/v3/pkg/application"
	"log"
	"os"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed frontend/dist
var assets embed.FS

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.

func main() {

	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	appService := &service.AppService{}
	app := application.New(application.Options{
		Name:        "SingBoxMu",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(appService),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	win := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "SingBoxMu",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
		Frameless:        true,
	})

	file, err := assets.ReadFile("frontend/dist/appicon.png")
	if err != nil {
		log.Fatal(err)
	}
	tray := app.NewSystemTray()
	tray.SetIcon(file)
	tray.SetLabel("SingBoxMu For Windows")
	menu := &application.Menu{}
	menu.Add("退出").OnClick(func(context *application.Context) {
		appService.StopCommand()
		app.Quit()
	})
	tray.SetMenu(menu)

	tray.OnClick(func() {
		if win.IsVisible() {
			win.Hide()
		} else {
			win.Show()
		}
	})

	// 如果执行参数为-hide
	if len(os.Args) > 1 {
		strings := os.Args[1:]
		// 如果包含字符
		if strings[0] == "-hide" {
			win.Hide()
		}
	}

	// Run the application. This blocks until the application has been exited.
	err = app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
