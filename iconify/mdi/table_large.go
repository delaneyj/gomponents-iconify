package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableLarge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 3h16a2 2 0 0 1 2 2v15a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2m0 4v3h4V7H4m6 0v3h4V7h-4m10 3V7h-4v3h4M4 12v3h4v-3H4m0 8h4v-3H4v3m6-8v3h4v-3h-4m0 8h4v-3h-4v3m10 0v-3h-4v3h4m0-8h-4v3h4v-3Z"/>`),
		g.Group(children),
	)
}