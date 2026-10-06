package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// version перезаписывается через -ldflags "-X main.version=..." во время сборки.
var version = "dev"

func main() {
	a := app.New()
	w := a.NewWindow("Hello GUI " + version)

	output := widget.NewLabel("Нажмите кнопку ниже")

	greetBtn := widget.NewButton("Поздороваться", func() {
		output.SetText(Greeting("GitHub"))
	})

	quitBtn := widget.NewButton("Выход", func() {
		a.Quit()
	})

	w.SetContent(container.NewVBox(
		widget.NewLabel("Hello from Go GUI! 🎨🐹"),
		widget.NewSeparator(),
		greetBtn,
		output,
		widget.NewSeparator(),
		widget.NewLabel(fmt.Sprintf("Version: %s", version)),
		widget.NewLabel(fmt.Sprintf("Sum 1..10 = %d", SumRange(1, 10))),
		widget.NewSeparator(),
		quitBtn,
	))

	w.ShowAndRun()
}
