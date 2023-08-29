package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarageOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.825 0-1.413-.588T2 20V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v16q0 .825-.588 1.413T20 22H4Zm0-2h16V4H4v16Zm5-6q-.425 0-.713-.288T8 13q0-.425.288-.713T9 12q.425 0 .713.288T10 13q0 .425-.288.713T9 14Zm6 0q-.425 0-.713-.288T14 13q0-.425.288-.713T15 12q.425 0 .713.288T16 13q0 .425-.288.713T15 14ZM5 11.1v6.6q0 .35.225.575t.575.225h.4q.35 0 .575-.225T7 17.7v-1.2h10v1.2q0 .35.225.575t.575.225h.4q.35 0 .575-.225T19 17.7v-6.6l-1.65-4.8q-.125-.35-.413-.575T16.3 5.5H7.7q-.35 0-.638.225T6.65 6.3L5 11.1Zm2.65-1.6l.7-2h7.3l.7 2h-8.7ZM4 4v16V4Zm3 10.5v-3h10v3H7Z"/>`),
		g.Group(children),
	)
}