package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Organisation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M17 15h2v2h-2v-2Zm2-4h-2v2h2v-2Z"/><path fill-rule="evenodd" d="M13 7h10v14H1V3h12v4ZM8 5h3v2H8V5Zm3 14v-2H8v2h3Zm0-4v-2H8v2h3Zm0-4V9H8v2h3Zm10 8V9h-8v2h2v2h-2v2h2v2h-2v2h8ZM3 19v-2h3v2H3Zm0-4h3v-2H3v2Zm3-4V9H3v2h3ZM3 7h3V5H3v2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}