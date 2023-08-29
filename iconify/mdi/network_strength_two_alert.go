package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkStrengthTwoAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 1L1 21h16v-2h-4v-7.2l6-6V9h2m-2 2v6h2v-6m-2 8v2h2v-2"/>`),
		g.Group(children),
	)
}