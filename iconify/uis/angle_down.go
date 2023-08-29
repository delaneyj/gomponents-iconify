package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.9 9.2c-.4-.4-1-.4-1.4 0L12 12.7L8.5 9.2c-.4-.4-1-.4-1.4 0s-.4 1 0 1.4l4.2 4.2c.2.2.4.3.7.3c.3 0 .5-.1.7-.3l4.2-4.2c.4-.4.4-1 0-1.4z"/>`),
		g.Group(children),
	)
}