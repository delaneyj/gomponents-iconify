package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.3 11.5c.2.2.4.3.7.3c.3 0 .5-.1.7-.3l3-3c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0L12 9.3L9.7 7c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l3 3.1zm3 1L12 14.8l-2.3-2.3c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l3 3c.2.2.4.3.7.3c.3 0 .5-.1.7-.3l3-3c.4-.4.4-1 0-1.4c-.4-.3-1-.3-1.4 0z"/>`),
		g.Group(children),
	)
}