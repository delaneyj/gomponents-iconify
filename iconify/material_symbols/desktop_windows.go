package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopWindows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21v-2h2v-2H4q-.825 0-1.413-.588T2 15V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v10q0 .825-.588 1.413T20 17h-6v2h2v2H8Z"/>`),
		g.Group(children),
	)
}