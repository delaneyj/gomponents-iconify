package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sorting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.7 17.8l-3-3c-.4-.4-1-.4-1.4 0l-3 3c-.4.4-.4 1 0 1.4c.4.4 1 .4 1.4 0l2.3-2.3l2.3 2.3c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4zm-4.4-7.6c.2.2.4.3.7.3c.3 0 .5-.1.7-.3l3-3c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0L12 8.1L9.7 5.8c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l3 3z"/>`),
		g.Group(children),
	)
}