package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerRightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.7 13.4c-.4-.4-1-.4-1.4 0l-2.9 2.9v-8c0-2.2-1.8-4-4-4H3c-.6 0-1 .4-1 1s.4 1 1 1h10.4c1.1 0 2 .9 2 2v8l-2.9-2.9c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l4.6 4.6c.2.2.4.3.7.3c.3 0 .5-.1.7-.3l4.6-4.6c.4-.4.4-1.1 0-1.4z"/>`),
		g.Group(children),
	)
}