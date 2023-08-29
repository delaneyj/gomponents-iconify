package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkunreadMailboxRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.825 0-1.413-.588T2 20V10q0-.825.588-1.413T4 8h2V3q0-.425.288-.713T7 2h6q.425 0 .713.288T14 3v2q0 .425-.288.713T13 6H8v7q0 .425.288.713T9 14q.425 0 .713-.288T10 13V8h10q.825 0 1.413.588T22 10v10q0 .825-.588 1.413T20 22H4Z"/>`),
		g.Group(children),
	)
}