package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 2H8C6.3 2 5 3.3 5 5v16c0 .2 0 .3.1.5c.3.5.9.6 1.4.4l5.5-3.2l5.5 3.2c.2.1.3.1.5.1c.6 0 1-.4 1-1V5c0-1.7-1.3-3-3-3z"/>`),
		g.Group(children),
	)
}