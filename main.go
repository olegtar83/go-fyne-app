package main

import (
	"giveout-emulator/tabs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Send to Queue")
	w.Resize(fyne.NewSize(300, 400))
	w.SetContent(container.NewAppTabs(
		container.NewTabItem("RabbitMq", tabs.GetRabbitTab()),
		container.NewTabItem("Azure Service Bus", widget.NewLabel("Will be implemented.")),
		container.NewTabItem("History", tabs.GetHistoryTab()),
	))
	w.ShowAndRun()
}
