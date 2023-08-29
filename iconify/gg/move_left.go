package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M19.071 17v-2h-8v2h8Zm0-8V7h-8v2h8Z"/><path d="M19.071 13v-2h-10V7.965l-4.071 4l4.071 4V13h10Z"/></g>`),
		g.Group(children),
	)
}