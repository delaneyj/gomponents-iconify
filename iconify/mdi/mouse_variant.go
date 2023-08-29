package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 7h-4V2.1A5 5 0 0 1 14 7M4 7a5 5 0 0 1 4-4.9V7H4m10 5a5 5 0 0 1-4 4.9V18a3 3 0 0 0 3 3a3 3 0 0 0 3-3v-5a4 4 0 0 1 4-4h2l-1 1l1 1h-2a2 2 0 0 0-2 2v5a5 5 0 0 1-5 5a5 5 0 0 1-5-5v-1.1A5 5 0 0 1 4 12V9h10v3Z"/>`),
		g.Group(children),
	)
}