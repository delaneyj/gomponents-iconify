package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsExpandDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V5Zm2 0h4v4h-4V5Z" clip-rule="evenodd"/><path d="M5 13H3v8h8v-2H6.414l5.364-5.364a1 1 0 0 0-1.414-1.414L5 17.586V13Z"/></g>`),
		g.Group(children),
	)
}