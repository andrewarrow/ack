package screen

import (
	"fmt"
	"infinity-dog/database"
	"log"
	"os"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var utc, _ = time.LoadLocation("UTC")
var services = widgets.NewList()
var messages = widgets.NewList()
var serviceItems = []database.Service{}
var offset = 0
var tab = "left"
var shutdown = false

func Setup() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	services.SelectedRowStyle.Fg = ui.ColorWhite
	services.SelectedRowStyle.Bg = ui.ColorMagenta
	services.TextStyle.Fg = ui.ColorWhite
	services.TextStyle.Bg = ui.ColorBlack
	serviceItems = database.ServicesByTotalBytes()
	for _, item := range serviceItems {
		services.Rows = append(services.Rows, fmt.Sprintf("% 11d %s", item.TotalBytes, item.Name))
	}
	messages.SelectedRowStyle.Fg = ui.ColorWhite
	messages.SelectedRowStyle.Bg = ui.ColorMagenta
	messages.TextStyle.Fg = ui.ColorWhite
	messages.TextStyle.Bg = ui.ColorBlack

	min, max := database.MinMaxDates()
	p2 := widgets.NewParagraph()
	p2.Text = fmt.Sprintf("[There](fg:blue,mod:bold) are [%d](fg:red) rows in sqlite. [From](fg:green) %.2f days ago to %.2f.", database.TotalRows(), min, max)
	//p2.SetRect(0, 5, 35, 10)
	p2.BorderStyle.Fg = ui.ColorYellow

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(0.40,
				ui.NewRow(0.8, services),
				ui.NewRow(0.2, p2)),
			ui.NewCol(0.59, messages),
		),
	)

	ui.Render(grid)
	uiEvents := ui.PollEvents()
	for {
		if shutdown {
			break
		}
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "j", "<Down>":
				if tab == "left" {
					services.ScrollDown()
				} else {
					messages.ScrollDown()
				}
			case "k", "<Up>":
				if tab == "left" {
					services.ScrollUp()
				} else {
					messages.ScrollUp()
				}
			case "<Tab>":
				if tab == "left" {
					tab = "right"
				} else {
					tab = "left"
				}
			case "<Left>":
				tab = "left"
			case "<Right>":
				tab = "right"
			case "<Enter>":
				offset = 0
				handleEnter()
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
			}
		}
		ui.Render(grid)
	}
}

func handleEnter() {
	if tab == "left" {
		serviceName := serviceItems[services.SelectedRow].Name
		items := database.MessagesFromService(serviceName)
		messages.Rows = []string{}
		utcNow := time.Now().In(utc).Unix()
		for _, item := range items {
			delta := float64(utcNow-item.LoggedAt) / 3600.0
			messages.Rows = append(messages.Rows, fmt.Sprintf("%.2f [%04d](fg:red) [%04d](fg:cyan) %s", delta, item.ExceptionLength, len(item.Both), item.BothTruncated(offset)))
		}
		messages.SelectedRow = 0
	} else if tab == "right" {

		serviceName := serviceItems[services.SelectedRow].Name
		items := database.MessagesFromService(serviceName)
		m := items[messages.SelectedRow]
		ui.Close()
		shutdown = true
		fmt.Println(m.Both)
		e := database.ExceptionById(m.Id)
		fmt.Println(e.Text)

		fmt.Println(m.Id)
		os.Exit(1)
	}
}
