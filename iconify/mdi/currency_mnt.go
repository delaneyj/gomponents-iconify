package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyMnt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 5v3.62l4-1.45v2.12l-4 1.45v1.76l4-1.43v2.13l-4 1.45V21h-2v-5.62l-4 1.46v-2.13l4-1.47v-1.77l-4 1.45V10.8l4-1.45V5H5V3h14v2h-6Z"/>`),
		g.Group(children),
	)
}