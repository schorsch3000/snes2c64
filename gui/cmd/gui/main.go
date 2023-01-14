package main

import (
	"snes2c64gui/cmd/gui/views"

	fyne "fyne.io/fyne/v2"

	"fyne.io/fyne/v2/app"
)

func main() {

	myApp := app.New()
	uploadView := views.NewUploadView()

	myWindow := myApp.NewWindow("Snes2C64")
	myWindow.Resize(fyne.NewSize(800, 400))
	myWindow.SetFixedSize(true)

	uploadView.Draw(myWindow)

	myWindow.Show()
	myApp.Run()
}