package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Power(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213 0v213h-42V0h42zm103 46q68 58 68 146q0 80-56 136t-136 56t-136-56T0 192q0-88 68-146l30 30q-55 45-55 116q0 62 43.5 105.5T192 341t105.5-43.5T341 192q0-71-55-115z"/>`),
		g.Group(children),
	)
}