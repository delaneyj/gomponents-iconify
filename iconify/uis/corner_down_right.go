package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.4 15.7L14.8 11c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l2.9 2.9h-7c-1.7 0-3-1.3-3-3V3c0-.6-.4-1-1-1s-1 .4-1 1v9.4c0 2.8 2.2 5 5 5h7l-2.9 2.9c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l4.6-4.6c.4-.4.4-1 0-1.4z"/>`),
		g.Group(children),
	)
}