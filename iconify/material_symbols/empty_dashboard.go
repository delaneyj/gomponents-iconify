package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmptyDashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18h5.5v-4H6v4Zm0-5h5.5V6H6v7Zm6.5 5H18v-7h-5.5v7Zm0-8H18V6h-5.5v4ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v2h2v2h-2v2h2v2h-2v2h2v2h-2v2q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}