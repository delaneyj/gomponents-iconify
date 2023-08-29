package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Messanger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feMessanger0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMessanger1" fill="currentColor"><path id="feMessanger2" d="m12.942 14.413l-2.56-2.66L5.45 14.48l5.406-5.736l2.56 2.66l4.93-2.727l-5.405 5.736ZM11.899 2C6.432 2 2 6.144 2 11.257c0 2.908 1.434 5.503 3.678 7.2V22l3.378-1.874c.9.252 1.855.388 2.843.388c5.468 0 9.9-4.145 9.9-9.257c0-5.113-4.432-9.257-9.9-9.257Z"/></g></g>`),
		g.Group(children),
	)
}