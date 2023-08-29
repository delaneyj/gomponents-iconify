package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-opacity=".3" d="M8 16h8V9h3v10H5V9h3v7Z"/><path d="M5 7h14V4H5v3Z"/></g>`),
		g.Group(children),
	)
}