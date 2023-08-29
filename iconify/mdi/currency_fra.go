package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyFra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 5V3H7v13H5v2h2v3h2v-3h4v-2H9v-3h8v-2H9V5h9Z"/>`),
		g.Group(children),
	)
}