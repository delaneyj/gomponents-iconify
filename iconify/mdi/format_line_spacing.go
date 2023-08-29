package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatLineSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 13h12v-2H10m0 8h12v-2H10m0-10h12V5H10M6 7h2.5L5 3.5L1.5 7H4v10H1.5L5 20.5L8.5 17H6V7Z"/>`),
		g.Group(children),
	)
}