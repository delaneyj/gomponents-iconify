package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyRub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21v-5H6v-1h1v-3H6v-1h1V4h7a4 4 0 0 1 4 4a4 4 0 0 1-4 4H8v3h6v1H8v5H7m1-10h6a3 3 0 0 0 3-3a3 3 0 0 0-3-3H8v6Z"/>`),
		g.Group(children),
	)
}