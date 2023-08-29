package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphaU(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 7v8a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V7h-2v8h-2V7H9Z"/>`),
		g.Group(children),
	)
}