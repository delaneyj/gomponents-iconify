package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.7 2c-.6 0-1 .4-1 1v10.4c0 1.1-.9 2-2 2h-8l2.9-2.9c.4-.4.4-1 0-1.4s-1-.4-1.4 0l-4.6 4.6c-.4.4-.4 1 0 1.4l4.6 4.6c.2.2.5.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4l-2.9-2.9h8c2.2 0 4-1.8 4-4V3c0-.6-.5-1-1-1z"/>`),
		g.Group(children),
	)
}