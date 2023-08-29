package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphaW(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 17a2 2 0 0 1-2-2V7h2v8h2V8h2v7h2V7h2v8a2 2 0 0 1-2 2H9Z"/>`),
		g.Group(children),
	)
}