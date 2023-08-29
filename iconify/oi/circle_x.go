package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4 0C1.79 0 0 1.79 0 4s1.79 4 4 4s4-1.79 4-4s-1.79-4-4-4zM2.5 1.78L4 3.28l1.5-1.5l.72.72L4.72 4l1.5 1.5l-.72.72L4 4.72l-1.5 1.5l-.72-.72L3.28 4l-1.5-1.5l.72-.72z"/>`),
		g.Group(children),
	)
}