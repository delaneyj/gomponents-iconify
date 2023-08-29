package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClinicalNotesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 16q-1.25 0-2.125-.875T14 13q0-1.25.875-2.125T17 10q1.25 0 2.125.875T20 13q0 1.25-.875 2.125T17 16Zm0-2q.425 0 .713-.288T18 13q0-.425-.288-.713T17 12q-.425 0-.713.288T16 13q0 .425.288.713T17 14Zm-6 9v-2.9q0-.525.25-.988t.7-.737q.8-.475 1.688-.788t1.812-.462L17 19l1.55-1.875q.925.15 1.8.463t1.675.787q.45.275.713.738T23 20.1V23H11Zm1.975-2h3.075l-1.35-1.65q-.45.125-.875.325t-.85.425v.9Zm4.975 0H21v-.9q-.4-.25-.825-.438t-.875-.312L17.95 21Zm-1.9 0Zm1.9 0ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v5q-.4-.5-.875-.95T19 8.45V5H5v14h4.15q-.075.275-.113.55T9 20.1v.9H5ZM7 9h7q.65-.5 1.425-.75T17 8V7H7v2Zm0 4h5q0-.525.112-1.025t.313-.975H7v2Zm0 4h3.45q.275-.225.588-.4t.637-.325V15H7v2Zm-2 2V5v3.425V8v11Zm12-6Z"/>`),
		g.Group(children),
	)
}