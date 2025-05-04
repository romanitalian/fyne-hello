package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create new app
	myApp := app.New()
	window := myApp.NewWindow("Hello App")

	// Create label that will show "Hello"
	helloLabel := widget.NewLabel("")
	helloLabel.Hide()

	// Create button
	button := widget.NewButton("Click me", func() {
		helloLabel.SetText("Hello")
		helloLabel.Show()
	})

	// Create vertical container with button and label
	content := container.NewVBox(
		container.NewCenter(button),
		container.NewCenter(helloLabel),
	)

	// Set content and show window
	window.SetContent(content)
	window.Resize(fyne.NewSize(200, 100))
	window.ShowAndRun()
}
