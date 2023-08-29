package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeVerticalCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 16h3v-2h14v2h3v2h-3v2H5v-2H2v-2m0-8h5v2h10V8h5V6h-5V4H7v2H2v2Z"/>`),
		g.Group(children),
	)
}