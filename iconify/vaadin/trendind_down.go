package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendindDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 14h-4l1.29-1.29L9 8.41l-3 3l-6-6V2.59l6 6l3-3l5.71 5.7L15.99 10l.01 4z"/>`),
		g.Group(children),
	)
}