package screen

import (
	"ack/files"
	"fmt"
	"log"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Transfer struct {
	source      *widgets.List
	transfer    *widgets.List
	destination *widgets.List
	text        string
}

func NewTransfer() *Transfer {
	t := Transfer{}
	t.source = widgets.NewList()
	t.transfer = widgets.NewList()
	t.destination = widgets.NewList()
	t.text = files.ReadFile("data/antony_and_cleopatra.txt")
	return &t
}

func (t *Transfer) Run() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	setList("Source", t.source)
	setList("", t.transfer)
	t.transfer.Border = false
	setList("Destination", t.destination)

	t.source.Rows = strings.Split(t.text, "\n")
	t.source.Title = fmt.Sprintf("Source %d", len(t.text))

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(0.33, t.source),
			ui.NewCol(0.33, t.transfer),
			ui.NewCol(0.33, t.destination),
		),
	)

	ui.Render(grid)
	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
			case "<Enter>":
				t.handleEnter()
			}
		}
		ui.Render(grid)
	}
}

func (t *Transfer) handleEnter() {
	//tcp.MAX_SEGMENT_SIZE
	t.transfer.Rows = append(t.transfer.Rows,
		fmt.Sprintf("Segment %d", len(t.transfer.Rows)))
}

func setList(title string, l *widgets.List) {
	l.SelectedRowStyle.Fg = ui.ColorWhite
	l.SelectedRowStyle.Bg = ui.ColorMagenta
	l.TextStyle.Fg = ui.ColorWhite
	l.TextStyle.Bg = ui.ColorBlack
	l.Title = title
}
