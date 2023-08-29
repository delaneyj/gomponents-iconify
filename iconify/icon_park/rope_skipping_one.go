package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RopeSkippingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke-miterlimit="2" d="M7 40.0001V22.0001C7 22.0001 8 8.00007 15 5.00008C22 2.00009 30 11.0001 30 20.0001C30 25.0001 28 29.0001 24 29.0001C20 29.0001 18 25.0001 18 20.0001C18 11.0001 26 2.00008 33 5.00008C40 8.00007 41 22.0001 41 22.0001V40.0001"/><path d="M4 31H10"/><path d="M38 31H44"/></g>`),
		g.Group(children),
	)
}