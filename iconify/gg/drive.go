package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M19 11a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-5 1a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z"/><path fill-rule="evenodd" d="M2 8a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2H2Zm20 2H2v4h20v-4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}