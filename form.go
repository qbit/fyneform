package fyneform

import (
	"fmt"
	"reflect"
	"strings"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type FormOpts interface {
	Titelize() bool
}

func MakeForm(st interface{}, opts FormOpts) (items []*widget.FormItem, err error) {
	stVal := reflect.ValueOf(st)
	strType := reflect.Indirect(stVal).Type()

	for i := 0; i < strType.NumField(); i++ {
		v := strType.Field(i)
		field := stVal.Elem().FieldByName(v.Name)

		if !field.IsValid() || !field.CanSet() {
			return nil, fmt.Errorf("can't set field: %q", v.Name)
		}

		switch field.Kind() {
		case reflect.Bool:
			if s, ok := field.Addr().Interface().(*bool); ok {
				boundBool := binding.BindBool(s)
				name := v.Tag.Get("fyneform")
				if name == "" {
					name = v.Tag.Get("json")
				}
				if opts.Titelize() {
					name = strings.Title(name)
				}

				w := widget.NewCheckWithData("", boundBool)
				items = append(items, widget.NewFormItem(name, w))
			}
		case reflect.String:
			if s, ok := field.Addr().Interface().(*string); ok {
				boundString := binding.BindString(s)

				w := widget.NewEntryWithData(boundString)
				name := v.Tag.Get("fyneform")
				if name == "" {
					name = v.Tag.Get("json")
				}
				if opts.Titelize() {
					name = strings.Title(name)
				}

				items = append(items, widget.NewFormItem(name, w))
			}
		}
	}
	return items, nil
}
