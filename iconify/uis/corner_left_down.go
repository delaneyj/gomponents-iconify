package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerLeftDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 4.3h-9.4c-2.8 0-5 2.2-5 5v7l-2.9-2.9c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l4.6 4.6c.2.2.4.3.7.3c.3 0 .5-.1.7-.3l4.6-4.6c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0l-2.9 2.9v-7c0-1.7 1.3-3 3-3H21c.6 0 1-.4 1-1s-.4-1-1-1z"/>`),
		g.Group(children),
	)
}