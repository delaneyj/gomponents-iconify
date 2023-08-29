package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertPolygon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m469 224l-52 59l8 79l-77 17l-41 68l-72-31l-73 31l-40-67l-77-18l7-79l-52-59l52-60l-7-78l77-17l40-68l73 31l72-31l41 68l77 17l-8 79zM256 331v-43h-43v43h43zm0-86V117h-43v128h43z"/>`),
		g.Group(children),
	)
}