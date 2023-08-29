package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.35 8.5L11 5h5V1h2v4h5l-1.38 13.79L11.35 8.5M1 21v1c0 .55.45 1 1 1h13c.55 0 1-.45 1-1v-1H1m20.9.9L2.1 2.1L.69 3.5l5.7 5.71C3.28 9.87 1 12 1 15h11.17l2 2H1v2h15v-.17l4.5 4.49l1.4-1.42Z"/>`),
		g.Group(children),
	)
}