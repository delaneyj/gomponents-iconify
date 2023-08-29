package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.775 18.925L14.85 12H20V8h-9.15l-4-4H20q.825 0 1.413.588T22 6v12q0 .25-.05.488t-.175.437ZM4 12h5.15l-4-4H4v4Zm16.45 11.3l-3.3-3.3H4q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4l2 2H3.15L.65 3.5l1.425-1.425l19.8 19.8L20.45 23.3Z"/>`),
		g.Group(children),
	)
}