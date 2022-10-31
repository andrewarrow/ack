package screen

import (
	"ack/files"
	"ack/tcp"
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type TransferWithOptions struct {
	source             *widgets.List
	transfer           *widgets.List
	destinationBuffer  *widgets.List
	destinationProgram *widgets.List
	text               string
	step               int
	segments           []string
}

func NewTransferWithOptions() *TransferWithOptions {
	t := TransferWithOptions{}
	t.source = widgets.NewList()
	t.transfer = widgets.NewList()
	t.destinationBuffer = widgets.NewList()
	t.destinationProgram = widgets.NewList()
	t.text = files.ReadFile("data/antony_and_cleopatra.txt")
	t.segments = tcp.BreakIntoSegments(t.text)
	return &t
}

func (t *TransferWithOptions) Run() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	setList("Source (server)", t.source)
	setList("", t.transfer)
	t.transfer.Border = false
	setList("Destination (buffer)", t.destinationBuffer)
	setList("Destination (program)", t.destinationProgram)

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(0.33, t.source),
			ui.NewCol(0.33, t.transfer),
			ui.NewCol(0.33,
				ui.NewRow(0.5, t.destinationBuffer),
				ui.NewRow(0.5, t.destinationProgram),
			),
		),
	)

	ui.Render(grid)
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Millisecond * 944).C
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
			}
		case <-ticker:
			t.advanceTransfer()
		}
		ui.Render(grid)
	}
}

func (t *TransferWithOptions) advanceTransfer() {
	t.destinationBuffer.Rows = append(t.destinationBuffer.Rows, "foo")
}
