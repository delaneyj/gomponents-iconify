package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toolbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 5.5h3a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-12a2 2 0 0 1 2-2h3a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3Zm-3-1h-4a1 1 0 0 0-1 1h6a1 1 0 0 0-1-1Zm6 3H4v2h16v-2Zm-16 12v-8h3v2h4v-2h2v2h4v-2h3v8H4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}