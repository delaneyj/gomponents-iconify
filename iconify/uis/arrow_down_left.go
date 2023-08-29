package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 16H9.4l8.3-8.3c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0L8 14.6V7c0-.6-.4-1-1-1s-1 .4-1 1v10c0 .1 0 .3.1.4c.1.2.3.4.5.5c.1.1.3.1.4.1h10c.6 0 1-.4 1-1s-.4-1-1-1z"/>`),
		g.Group(children),
	)
}