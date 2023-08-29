package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 6c-.6 0-1 .4-1 1v7.6L7.7 6.3c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l8.3 8.3H7c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1V7c0-.6-.4-1-1-1z"/>`),
		g.Group(children),
	)
}