package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.4 8H17c.6 0 1-.4 1-1s-.4-1-1-1H7c-.6 0-1 .4-1 1v10c0 .6.4 1 1 1s1-.4 1-1V9.4l8.3 8.3c.4.4 1 .4 1.4 0c.4-.4.4-1 0-1.4L9.4 8z"/>`),
		g.Group(children),
	)
}