package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotNet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 15a1 1 0 0 1 1 1a1 1 0 0 1-1 1a1 1 0 0 1-1-1a1 1 0 0 1 1-1m19 2h-2V9h-2V7h6v2h-2v8M16 7v2h-2v2h2v2h-2v2h2v2h-4V7h4m-5 0v10H9l-3-6v6H4V7h2l3 6V7h2Z"/>`),
		g.Group(children),
	)
}