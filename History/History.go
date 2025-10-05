package History

//
//import (
//	"curGUI/Currencies/JSON"
//
//	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/canvas"
//	"fyne.io/fyne/v2/container"
//	"fyne.io/fyne/v2/widget"
//)
//
//func HistoryWIN(History []JSON.JSON) (fyne.Container, widget.Entry, widget.Entry, widget.Entry, fyne.Container) {
//	IN := widget.NewEntry()
//	OUT := widget.NewEntry()
//	Currency := widget.NewEntry()
//
//	IN.Resize(fyne.NewSize(230, 60))
//	IN.Move(fyne.NewPos(35, 260))
//	OUT.Resize(fyne.NewSize(230, 65))
//	OUT.Move(fyne.NewPos(35, 350))
//	Currency.Resize(fyne.NewSize(230, 55))
//	Currency.Move(fyne.NewPos(35, 445))
//
//	History2 := container.NewVBox()
//
//	background2, _ := fyne.LoadResourceFromPath("His.png")
//	Back2 := canvas.NewImageFromResource(background2)
//	Back2.FillMode = canvas.ImageFillContain
//	Back2.Resize(fyne.NewSize(300, 560))
//
//	scroll := container.NewScroll(History2)
//	return container.NewWithoutLayout(Back2, scroll, IN, OUT, Currency), IN, OUT, Currency, History2
//}
