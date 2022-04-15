package tabs

import (
	"encoding/json"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var stringData = `[{"date":"25.04.1987","pass":"8787675656545","type":"rabbit"},{"date":"12.04.2013","pass":"54934543","type":"azure sb"}]`

func GetHistoryTab() *fyne.Container {
	var data []Message

	json.Unmarshal([]byte(stringData), &data)

	var bindings []binding.DataMap
	for _, messages := range data {
		bindings = append(bindings, binding.BindStruct(&messages))
	}

	hello := widget.NewLabel("History of sendings.")
	table := widget.NewTable(
		func() (int, int) {
			return len(bindings), 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			row, col := i.Row, i.Col
			var item binding.DataItem
			switch col {
			case 0:
				item, _ = bindings[row].GetItem("Date")
			case 1:
				item, _ = bindings[row].GetItem("Pass")
			case 2:
				item, _ = bindings[row].GetItem("Type")
			}
			o.(*widget.Label).Bind(item.(binding.String))
		},
	)

	// table.MinSize().Min(fyne.NewSize(380, 450))
	// table.MinSize().Max(fyne.NewSize(380, 450))
	container := container.NewVBox(hello, table)
	return container
}

type Message struct {
	Date string `json:"date,omitempty"`
	Pass string `json:"pass,omitempty"`
	Type string `json:"type,omitempty"`
}
