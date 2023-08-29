package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 5h13a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3m0 1a2 2 0 0 0-2 2v1h17V8a2 2 0 0 0-2-2H5M3 17a2 2 0 0 0 2 2h13a2 2 0 0 0 2-2v-5H3v5m2-1h4v1H5v-1m6 0h3v1h-3v-1m-8-6v1h17v-1H3Z"/>`),
		g.Group(children),
	)
}