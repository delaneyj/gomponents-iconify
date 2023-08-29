package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Details(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.38 6h11.25L12 16L6.38 6M3 4l9 16l9-16H3Z"/>`),
		g.Group(children),
	)
}