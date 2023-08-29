package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SetMerge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7v2h5V7H2m10 2v2H9v2h3v2l3-3l-3-3m5 0v6h5V9h-5M2 11v2h5v-2H2m0 4v2h5v-2H2Z"/>`),
		g.Group(children),
	)
}