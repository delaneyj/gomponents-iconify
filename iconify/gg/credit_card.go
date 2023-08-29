package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4 9a1 1 0 0 1 1-1h4a1 1 0 0 1 0 2H5a1 1 0 0 1-1-1Z"/><path fill-rule="evenodd" d="M4 3a4 4 0 0 0-4 4v10a4 4 0 0 0 4 4h16a4 4 0 0 0 4-4V7a4 4 0 0 0-4-4H4Zm16 2H4a2 2 0 0 0-2 2v7h20V7a2 2 0 0 0-2-2Zm2 11H2v1a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-1Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}