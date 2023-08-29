package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SizeS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 7c-1.1 0-2 .9-2 2v2a2 2 0 0 0 2 2h2v2H9v2h4c1.11 0 2-.89 2-2v-2a2 2 0 0 0-2-2h-2V9h4V7h-4Z"/>`),
		g.Group(children),
	)
}