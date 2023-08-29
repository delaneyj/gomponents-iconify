package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Duplicate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M19 5H7V3h14v14h-2V5ZM9 13v-2h2v2h2v2h-2v2H9v-2H7v-2h2Z"/><path fill-rule="evenodd" d="M3 7h14v14H3V7Zm2 2h10v10H5V9Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}