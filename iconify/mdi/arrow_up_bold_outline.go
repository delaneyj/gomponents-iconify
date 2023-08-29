package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpBoldOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 13v8H8v-8H2L12 3l10 10h-6m-9-2h3v8h4v-8h3l-5-5l-5 5Z"/>`),
		g.Group(children),
	)
}