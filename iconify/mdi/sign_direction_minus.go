package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignDirectionMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 9.5L18 12h-5v10H9a2 2 0 0 1 2-2v-8H3.5L6 9.5L3.5 7H11V3l1-1l1 1v4h5m5 11h-8v2h8Z"/>`),
		g.Group(children),
	)
}