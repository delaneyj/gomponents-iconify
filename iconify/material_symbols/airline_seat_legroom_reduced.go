package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatLegroomReduced(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 21v-3l1-4H8q-.825 0-1.413-.588T6 12V3h6v6h5q.825 0 1.413.588T19 11l-2 7h1.45q.675 0 1.113.413T20 19.5q0 .675-.413 1.088T18.5 21H14Zm-2-4H5q-.825 0-1.413-.588T3 15V3h2v12h7v2Z"/>`),
		g.Group(children),
	)
}