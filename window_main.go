package main

//go build -ldflags -H=windowsgui
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"strconv"
)

func main() {

	apps := app.New()
	win := apps.NewWindow("Golang Macros BF4")

	win.Resize(fyne.Size{
		Width:  250,
		Height: 150,
	})

	res, err := LoadResourceFromPath(JPGPath)
	if err != nil {
		log.Println(err)
	}
	win.SetIcon(res)

	go Cheat_Clicker()
	go Cheat_Burstmacro()
	go Kill_Process()
	//go GetAllTokens()

	label1 := widget.NewLabel("Vehicle Clicker: Ctrl + Y")
	label2 := widget.NewLabel("Burst Macros: Ctrl + B")
	label3 := widget.NewLabel("Fast Quit: Ctrl + K")

	tag := canvas.NewText("by Amobus", color.Black)
	tag.TextSize = 11
	tag.Alignment = fyne.TextAlignCenter

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Click delay in ms:")
	entry.SetText("5")

	entry.OnChanged = func(text string) {
		iDelay, err = strconv.Atoi(text)
		if err != nil {
			dialog.ShowError(err, win)
		}
	}

	win.SetContent(container.NewVBox(label1,
		entry, label2, label3, tag, // label4,
	),
	)

	win.CenterOnScreen()
	win.ShowAndRun()
}
