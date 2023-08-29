package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsExpandUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M13 5V3h8v8h-2V6.414l-5.364 5.364a1 1 0 0 1-1.414-1.414L17.586 5H13Z"/><path fill-rule="evenodd" d="M5 13a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2H5Zm0 2v4h4v-4H5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}