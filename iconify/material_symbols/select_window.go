package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectWindow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.825 0-1.413-.588T2 20v-9q0-.825.588-1.413T4 9h2V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v9q0 .825-.588 1.413T20 15h-2v5q0 .825-.588 1.413T16 22H4Zm0-2h12v-7H4v7Zm14-7h2V6H8v3h8q.825 0 1.413.588T18 11v2Z"/>`),
		g.Group(children),
	)
}