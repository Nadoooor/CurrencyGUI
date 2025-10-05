package main

import (
	"curGUI/Currencies"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	var selected string
	var selected2 string
	w := a.NewWindow("CurrencyGUI")
	background, _ := fyne.LoadResourceFromPath("Back.png")
	Back := canvas.NewImageFromResource(background)
	Back.FillMode = canvas.ImageFillContain
	Back.Resize(fyne.NewSize(300, 560))

	currencies := []string{"EGP", "USD", "EUR", "GBP", "JPY", "CAD", "CHF", "CNY", "INR", "HKD"}

	drop1 := widget.NewSelect(currencies, func(s string) {
		selected = s
	})

	currencies2 := []string{"EGP", "USD", "EUR", "GBP", "JPY", "CAD", "CHF", "CNY", "INR", "HKD"}
	drop2 := widget.NewSelect(currencies2, func(s string) {
		selected2 = s
	})
	input := widget.NewMultiLineEntry()
	output := widget.NewMultiLineEntry()

	Convert := widget.NewButton("Convert", func() {
		amount, _ := strconv.ParseFloat(input.Text, 64)
		result := Currencies.Conv(selected2, selected, amount)
		output.SetText(strconv.FormatFloat(result, 'f', -1, 64))

	})

	// UI resize and moving
	drop2.Resize(fyne.NewSize(95, 40))
	drop2.Move(fyne.NewPos(35, 180))
	drop1.Resize(fyne.NewSize(95, 40))
	drop1.Move(fyne.NewPos(170, 180))
	Convert.Resize(fyne.NewSize(230, 50))
	Convert.Move(fyne.NewPos(35, 455))
	input.Resize(fyne.NewSize(230, 35))
	input.Move(fyne.NewPos(35, 108))
	output.Resize(fyne.NewSize(230, 150))
	output.Move(fyne.NewPos(35, 260))

	maincon := container.NewWithoutLayout(Back, drop1, drop2, Convert, input, output)
	//Hiscon := container.NewWithoutLayout(Back2, IN, OUT, Currency)
	taps := container.NewAppTabs(
		container.NewTabItem("Convertor", maincon),
		//container.NewTabItem("History", Hiscon),
	)
	w.SetContent(taps)
	w.Resize(fyne.NewSize(300, 600))
	w.SetFixedSize(true)
	w.ShowAndRun()
}
