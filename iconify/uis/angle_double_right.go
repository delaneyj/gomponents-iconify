package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 8.3c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4L9.3 12L7 14.3c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l3-3c.4-.4.4-1 0-1.4l-2.9-3zm8.5 3l-3-3c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l2.3 2.3l-2.3 2.3c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3l3-3c.3-.4.3-1 0-1.4z"/>`),
		g.Group(children),
	)
}