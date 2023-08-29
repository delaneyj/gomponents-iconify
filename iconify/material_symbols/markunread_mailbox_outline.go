package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkunreadMailboxOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14v-4H4v10h16V10H10V8h10q.825 0 1.413.588T22 10v10q0 .825-.588 1.413T20 22H4q-.825 0-1.413-.588T2 20V10q0-.825.588-1.413T4 8h2V2h8v4H8v8H6Zm-2-4v4v-4v10v-10Z"/>`),
		g.Group(children),
	)
}