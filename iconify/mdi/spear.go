package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 9h.41l-13 13L2 20.59l13-13V9h1m0-5v4h4l2-6l-6 2Z"/>`),
		g.Group(children),
	)
}