package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.1 21.9l-17-17q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275ZM14 11.15l-2-2V5q0-.825.588-1.413T14 3h2q.825 0 1.413.588T18 5q0 .825-.588 1.413T16 7h-2v4.15ZM10 21q-1.65 0-2.825-1.175T6 17q0-1.65 1.175-2.825T10 13q.575 0 1.063.138t.937.412V12l2 2v3q0 1.65-1.175 2.825T10 21Z"/>`),
		g.Group(children),
	)
}