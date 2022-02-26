package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"suah.dev/fyneform"
)

type test struct {
	First  string `json:"first"`
	Last   string `fyneform:"Last Name" json:"last"`
	Pass   string `fyneform:"Passphrase" fynetype:"password" json:"pass"`
	Enable bool   `fyneform:"Enable?"`
}

type opts struct{}

func (o *opts) Titelize() bool {
	return false
}

func (o *opts) EntryOverride(tags reflect.StructTag, bindType interface{}) (bool, *widget.Entry) {
	return false, nil
}

func main() {
	f := &test{}
	o := &opts{}

	a := app.NewWithID("dev.suah.fyne_struct")
	bg := canvas.NewRectangle(color.Gray{Y: 0x16})
	w := a.NewWindow("Fyne Struct")
	w.SetContent(container.NewMax(bg))
	w.Resize(fyne.NewSize(300, 400))
	items, err := fyneform.MakeForm(f, o)
	if err != nil {
		log.Fatal(err)
	}

	dialog.ShowForm("Fyne Struct Test", "Ok", "Cancel", items,
		func(ok bool) {
			if ok {
				fmt.Printf("%#v\n", f)
			}
			os.Exit(0)
		}, w)
	w.ShowAndRun()
}
