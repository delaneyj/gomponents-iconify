package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottomRightClickSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h16V3h2v18H3Zm12-4q-.825 0-1.413-.588T13 15q0-.825.588-1.413T15 13q.825 0 1.413.588T17 15q0 .825-.588 1.413T15 17ZM5 12v-2h3.6L3 4.4L4.4 3L10 8.6V5h2v7H5Z"/>`),
		g.Group(children),
	)
}