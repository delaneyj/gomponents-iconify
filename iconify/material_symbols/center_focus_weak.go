package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CenterFocusWeak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19v-4h2v4h4v2H5Zm10 0v-2h4v-4h2v4q0 .825-.588 1.413T19 21h-4ZM3 9V5q0-.825.588-1.413T5 3h4v2H5v4H3Zm16 0V5h-4V3h4q.825 0 1.413.588T21 5v4h-2Zm-7 7q-1.65 0-2.825-1.175T8 12q0-1.65 1.175-2.825T12 8q1.65 0 2.825 1.175T16 12q0 1.65-1.175 2.825T12 16Z"/>`),
		g.Group(children),
	)
}