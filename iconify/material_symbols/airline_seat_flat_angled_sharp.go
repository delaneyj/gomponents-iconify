package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatFlatAngledSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.25 17.175l-12.225-4.45l2.4-6.575L23.65 10.6l-2.4 6.575ZM20.225 20l-18.8-6.85l.675-1.875l18.8 6.85L20.225 20Zm-13.9-8.45q-1.25 0-2.125-.875T3.325 8.55q0-1.25.875-2.125t2.125-.875q1.25 0 2.125.875t.875 2.125q0 1.25-.875 2.125t-2.125.875Z"/>`),
		g.Group(children),
	)
}