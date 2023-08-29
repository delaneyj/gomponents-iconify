package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiveGMobiledataBadge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q-.825 0-1.413-.588T1 19V5q0-.825.588-1.413T3 3h18q.825 0 1.413.588T23 5v14q0 .825-.588 1.413T21 21H3Zm16-10h-3v2h1v2h-3V9h5q0-.825-.588-1.413T17 7h-3q-.825 0-1.413.588T12 9v6q0 .825.588 1.413T14 17h3q.825 0 1.413-.588T19 15v-4ZM5 17h4q.825 0 1.413-.588T11 15v-2q0-.825-.588-1.413T9 11H7V9h4V7H5v6h4v2H5v2Z"/>`),
		g.Group(children),
	)
}