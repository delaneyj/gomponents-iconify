package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5 17v-2h8v2H5Zm0-8V7h8v2H5Z"/><path d="M5 13v-2h10V7.965l4.071 4l-4.071 4V13H5Z"/></g>`),
		g.Group(children),
	)
}