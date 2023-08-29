package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiveG(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17v-2h5v-2H3V7h7v2H5v2h3q.825 0 1.413.588T10 13v2q0 .825-.588 1.413T8 17H3Zm18-6v4q0 .825-.588 1.413T19 17h-5q-.825 0-1.413-.588T12 15V9q0-.825.588-1.413T14 7h5q.825 0 1.413.588T21 9h-7v6h5v-2h-2.5v-2H21Z"/>`),
		g.Group(children),
	)
}