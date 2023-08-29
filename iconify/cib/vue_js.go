package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VueJs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.306 4.019H19.5L16 9.556l-3-5.537H2L16 28L30 4.019zm-18.825 2h3.363L16 18.406l7.15-12.387h3.363L16.001 24.031z"/>`),
		g.Group(children),
	)
}