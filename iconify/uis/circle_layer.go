package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleLayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 12c-2.8 0-5 2.2-5 5s2.2 5 5 5s5-2.2 5-5s-2.2-5-5-5zm4-5c-2.4 0-4.6 1.5-5.5 3.7c3.5-.9 7 1.3 7.8 4.7c.3 1 .3 2.1 0 3.1c3.1-1.3 4.5-4.8 3.2-7.8C15.6 8.5 13.4 7 11 7zm10.2-1.2C20.1 3.5 17.6 2 15 2S9.9 3.5 8.8 5.8c4-1.2 8.2 1 9.4 4.9c.5 1.5.5 3 0 4.5c3.4-1.7 4.8-5.9 3-9.4z"/>`),
		g.Group(children),
	)
}