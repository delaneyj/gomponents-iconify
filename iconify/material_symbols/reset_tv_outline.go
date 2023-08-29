package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResetTvOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21v-2H4q-.825 0-1.413-.588T2 17V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v3h-2V5H4v12h16v-5h-7.2l1.85 1.85l-1.4 1.4L9 11l4.25-4.25l1.4 1.4L12.8 10H20q.825 0 1.413.588T22 12v5q0 .825-.588 1.413T20 19h-4v2H8Zm5-10Z"/>`),
		g.Group(children),
	)
}