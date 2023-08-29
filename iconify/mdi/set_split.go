package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SetSplit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 7v2h5V7h-5M2 9v6h5V9H2m10 0v2H9v2h3v2l3-3l-3-3m5 2v2h5v-2h-5m0 4v2h5v-2h-5Z"/>`),
		g.Group(children),
	)
}