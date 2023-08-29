package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditBlackPoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8 12a4 4 0 1 1 8 0a4 4 0 0 1-8 0Zm4 1a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/><path d="M3 12a9 9 0 1 1 18 0a9 9 0 0 1-18 0Zm9 7a7 7 0 1 1 0-14a7 7 0 0 1 0 14Z"/></g>`),
		g.Group(children),
	)
}