package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipNextOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 18l8.5-6L6 6m2 3.86L11.03 12L8 14.14M16 6h2v12h-2"/>`),
		g.Group(children),
	)
}