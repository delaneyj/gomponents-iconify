package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.7 6.6h-7l2.9-2.9c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0L4.6 6.9c-.4.4-.4 1 0 1.4L9.2 13c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4l-2.9-3h7c1.7 0 3 1.3 3 3V21c0 .6.4 1 1 1s1-.4 1-1v-9.4c0-2.7-2.3-5-5-5z"/>`),
		g.Group(children),
	)
}