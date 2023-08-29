package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopScreenShareRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.2 21H2q-.425 0-.713-.288T1 20q0-.425.288-.713T2 19h14.175l-1-1H4q-.825 0-1.413-.588T2 16V4.85l-.625-.65Q1.1 3.9 1.1 3.5t.3-.7q.275-.275.7-.275t.7.275l18.4 18.4q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L18.2 21ZM10 12.85L8.175 11q-.125.225-.15.475T8 12v1q0 .425.288.713T9 14q.425 0 .713-.288T10 13v-.15Zm10.7 5l-6.275-6.275l1.225-1.225q.15-.15.15-.35t-.15-.35l-2.225-2.225q-.125-.125-.275-.062T13 7.6V9h-1.15l-6-6H20q.825 0 1.413.588T22 5v11q0 .65-.363 1.137t-.937.713Z"/>`),
		g.Group(children),
	)
}