package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipPreviousOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 6h2v12H6m3.5-6l8.5 6V6m-2 8.14L12.97 12L16 9.86v4.28Z"/>`),
		g.Group(children),
	)
}