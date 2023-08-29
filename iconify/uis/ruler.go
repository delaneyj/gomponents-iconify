package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22.6 7.1l-5.7-5.7c-.4-.4-1-.4-1.4 0L1.4 15.5c-.4.4-.4 1 0 1.4l5.7 5.7c.4.4 1 .4 1.4 0l2.1-2.1L7.1 17c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0l3.5 3.5l1.4-1.4l-2.1-2.1c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0l2.1 2.1l1.4-1.4l-3.5-3.5c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0l3.5 3.5l1.4-1.4l-2.1-2.1c-.4-.4-.4-1 0-1.4s1-.4 1.4 0l2.1 2.1l2.1-2.1c.5-.5.5-1.2.1-1.5z"/>`),
		g.Group(children),
	)
}