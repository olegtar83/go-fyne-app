package tabs

import (
	"giveout-emulator/models"
	"giveout-emulator/rabbit"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GetRabbitTab() *fyne.Container {
	hello := widget.NewLabel("Send to RabbitMQ!")
	button1 := widget.NewButton("000000000000056223", func() {
		pass := &models.PassModel{
			CardType:    "3",
			PassCode:    "000000000000056223",
			Serial:      "5506",
			Workstation: "wpet-w10-p00320",
		}
		rabbit.PublishMessage(pass)
		hello.SetText("Sent: " + pass.PassCode)
	})
	button2 := widget.NewButton("000000000000141136", func() {
		pass := &models.PassModel{
			CardType:    "3",
			PassCode:    "000000000000141136",
			Serial:      "5506",
			Workstation: "wpet-w10-p00320",
		}
		rabbit.PublishMessage(pass)
		hello.SetText("Sent: " + pass.PassCode)
	})
	button3 := widget.NewButton("000000000014655060", func() {
		pass := &models.PassModel{
			CardType:    "3",
			PassCode:    "000000000014655060",
			Serial:      "5506",
			Workstation: "wpet-w10-p00320",
		}
		rabbit.PublishMessage(pass)
		hello.SetText("Sent: " + pass.PassCode)
	})
	button4 := widget.NewButton("36096178920526596", func() {
		pass := &models.PassModel{
			CardType:    "3",
			PassCode:    "36096178920526596",
			Serial:      "5506",
			Workstation: "wpet-w10-p00320",
		}
		rabbit.PublishMessage(pass)
		hello.SetText("Sent: " + pass.PassCode)
	})
	button5 := widget.NewButton("00000000000132697", func() {
		pass := &models.PassModel{
			CardType:    "3",
			PassCode:    "00000000000132697",
			Serial:      "5506",
			Workstation: "wpet-w10-p00320",
		}
		rabbit.PublishMessage(pass)
		hello.SetText("Sent: " + pass.PassCode)
	})
	return container.NewVBox(hello, button1, button2, button3, button4, button5)
}
