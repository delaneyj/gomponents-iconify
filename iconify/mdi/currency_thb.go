package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyThb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.1 11.6c.6-.7.9-1.6.9-2.6c0-1.9-1.3-3.4-3-3.9L13 5V3h-2v2H7v14h4v2h2v-2h1c2.2 0 4-1.8 4-4c0-1.5-.8-2.7-1.9-3.4M15 9c0 1.1-.9 2-2 2V7c1.1 0 2 .9 2 2M9 7h2v4H9V7m0 10v-4h2v4H9m5 0h-1v-4h1c1.1 0 2 .9 2 2s-.9 2-2 2Z"/>`),
		g.Group(children),
	)
}