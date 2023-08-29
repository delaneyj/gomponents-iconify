package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreaterThanOrEqual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 2.27L20 10.14L6.5 18l-1-1.73l10.53-6.13L5.5 4l1-1.73M20 20v2H5v-2h15Z"/>`),
		g.Group(children),
	)
}