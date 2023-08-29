package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpDownBoldOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 10h6L12 0L2 10h6v4H2l10 10l10-10h-6v-4m-2 6h3l-5 5l-5-5h3V8H7l5-5l5 5h-3v8Z"/>`),
		g.Group(children),
	)
}