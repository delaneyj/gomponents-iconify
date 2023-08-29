package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenJam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21v-2h7v-7.2l1.6 1.6L16 12l-4-4l-4 4l1.4 1.4l1.6-1.6V16H4q-.825 0-1.413-.588T2 14V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v9q0 .825-.588 1.413T20 16h-5v3h3v2H6Z"/>`),
		g.Group(children),
	)
}