package screen

import (
	"fmt"
	"log"
	"sync"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var mutex sync.Mutex

type TransferWithOptions struct {
	source             *widgets.List
	transfer           *widgets.List
	destinationBuffer  *widgets.List
	destinationProgram *widgets.List
	step               int
}

func NewTransferWithOptions() *TransferWithOptions {
	t := TransferWithOptions{}
	t.source = widgets.NewList()
	t.transfer = widgets.NewList()
	t.destinationBuffer = widgets.NewList()
	t.destinationProgram = widgets.NewList()
	t.step = 301
	return &t
}

func (t *TransferWithOptions) Run() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	setList("Source (server)", t.source)
	t.source.Rows = append(t.source.Rows, "102.7 MB")
	t.source.Rows = append(t.source.Rows, "102,729,562 bytes")
	t.source.Rows = append(t.source.Rows, "102,729,562 / 512")
	t.source.Rows = append(t.source.Rows, "200,643.67578125")
	t.source.Rows = append(t.source.Rows, "200,644 segments")
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
	ticker1 := time.NewTicker(time.Millisecond * 944).C
	ticker2 := time.NewTicker(time.Millisecond * 3944).C
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
		case <-ticker1:
			t.advanceTransfer()
		case <-ticker2:
			t.readBuffer()
		}
		ui.Render(grid)
	}
}

func (t *TransferWithOptions) advanceTransfer() {
	mutex.Lock()
	defer mutex.Unlock()
	if len(t.destinationBuffer.Rows) == 10 {
		return
	}
	t.destinationBuffer.Rows = append([]string{fmt.Sprintf("Seg %03d", t.step)}, t.destinationBuffer.Rows...)
	fullString := ""
	if len(t.destinationBuffer.Rows) == 10 {
		fullString = "FULL"
	}
	t.destinationBuffer.Title = fmt.Sprintf("Size %d %s", len(t.destinationBuffer.Rows), fullString)
	t.step++
}

func (t *TransferWithOptions) readBuffer() {
	mutex.Lock()
	defer mutex.Unlock()
	item := ""
	if len(t.destinationBuffer.Rows) > 0 {
		item = t.destinationBuffer.Rows[len(t.destinationBuffer.Rows)-1]
		t.destinationBuffer.Rows = t.destinationBuffer.Rows[0 : len(t.destinationBuffer.Rows)-1]
		t.destinationBuffer.Title = fmt.Sprintf("Size %d", len(t.destinationBuffer.Rows))
	}
	if item != "" {
		t.destinationProgram.Rows = append([]string{item}, t.destinationProgram.Rows...)
	}
	t.destinationProgram.Title = fmt.Sprintf("Size %d", len(t.destinationProgram.Rows))
}
