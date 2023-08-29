package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Studio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M17 13h-4v4h4v-4Z"/><path fill-rule="evenodd" d="M3 3h18v18H3V3Zm2 2h14v14H5V5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}