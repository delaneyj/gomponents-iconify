package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CenterFocusStrong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19v-4h2v4h4v2H5Zm10 0v-2h4v-4h2v4q0 .825-.588 1.413T19 21h-4ZM3 9V5q0-.825.588-1.413T5 3h4v2H5v4H3Zm16 0V5h-4V3h4q.825 0 1.413.588T21 5v4h-2Zm-7 8q-2.075 0-3.538-1.463T7 12q0-2.075 1.463-3.538T12 7q2.075 0 3.538 1.463T17 12q0 2.075-1.463 3.538T12 17Z"/>`),
		g.Group(children),
	)
}