package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m0 11.5l4-4l8.5 8.5l3.5-3.5L7.5 4l4-4H0v11.5z"/>`),
		g.Group(children),
	)
}