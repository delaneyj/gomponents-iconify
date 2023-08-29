package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarShiftPattern(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 5H4V2h4v3M4 22h4v-3H4v3M14 2h-4v3h4V2m-4 20h4v-3h-4v3m6-20v3h4V2h-4m1 9h-4V7h-2v4H7V7H5v10h2v-4h4v4h2v-4h6V7h-2v4Z"/>`),
		g.Group(children),
	)
}