package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14v2H4q-.825 0-1.413-.588T2 14V4q0-.825.588-1.413T4 2h10q.825 0 1.413.588T16 4v2h-2V4H4v10h2Zm4 8q-.825 0-1.413-.588T8 20V10q0-.825.588-1.413T10 8h10q.825 0 1.413.588T22 10v10q0 .825-.588 1.413T20 22H10Z"/>`),
		g.Group(children),
	)
}