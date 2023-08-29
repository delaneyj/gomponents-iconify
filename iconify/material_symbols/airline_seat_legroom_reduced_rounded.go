package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatLegroomReducedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 21q-.8 0-1.363-.65t-.362-1.425L15 14H8q-.825 0-1.413-.588T6 12V3h6v6h5q.825 0 1.413.588T19 11l-2 7h1.45q.675 0 1.113.413T20 19.5q0 .675-.413 1.088T18.5 21h-3ZM5 17q-.825 0-1.413-.588T3 15V4q0-.425.288-.713T4 3q.425 0 .713.288T5 4v11h6q.425 0 .713.288T12 16q0 .425-.288.713T11 17H5Z"/>`),
		g.Group(children),
	)
}