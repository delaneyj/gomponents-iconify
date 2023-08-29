package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpertsExchange(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9.708 1.198H0L11.146 16L0 30.802h9.708L20.854 16zm22.292 0h-9.708l-3.063 4.083l4.849 6.427zM19.229 26.734l3.063 4.068H32l-7.922-10.51z"/>`),
		g.Group(children),
	)
}