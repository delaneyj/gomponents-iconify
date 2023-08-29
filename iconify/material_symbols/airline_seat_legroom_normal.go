package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatLegroomNormal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 21v-7H8q-.825 0-1.413-.588T6 12V3h6v6h5q.825 0 1.413.588T19 11v7h1.5q.65 0 1.075.425T22 19.5q0 .65-.425 1.075T20.5 21H16Zm-2-4H5q-.825 0-1.413-.588T3 15V3h2v12h9v2Z"/>`),
		g.Group(children),
	)
}