package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 21q-1.65 0-2.825-1.175T6 17q0-1.65 1.175-2.825T10 13q.575 0 1.063.138t.937.412V5q0-.825.588-1.413T14 3h2q.825 0 1.413.588T18 5q0 .825-.588 1.413T16 7h-2v10q0 1.65-1.175 2.825T10 21Z"/>`),
		g.Group(children),
	)
}