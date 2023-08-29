package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.4 6.9l-4.6-4.6c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l2.9 2.9h-8c-2.2 0-4 1.8-4 4V21c0 .6.4 1 1 1s1-.4 1-1V10.6c0-1.1.9-2 2-2h8l-2.9 2.9c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l4.6-4.6c.4-.4.4-1 0-1.4z"/>`),
		g.Group(children),
	)
}