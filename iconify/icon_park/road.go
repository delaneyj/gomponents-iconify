package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Road(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M11 6L6 42"/><path d="M37 6L41.9644 41.9552"/><path d="M24 6V12"/><path d="M24 35V42"/><path d="M24 20V27"/></g>`),
		g.Group(children),
	)
}