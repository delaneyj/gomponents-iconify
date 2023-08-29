package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkStrengthOneAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 1L1 21h16v-2h-6v-5.2l8-8V9h2m-2 2v6h2v-6m-2 8v2h2v-2"/>`),
		g.Group(children),
	)
}