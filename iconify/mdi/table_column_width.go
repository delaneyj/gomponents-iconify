package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableColumnWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 8h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V10a2 2 0 0 1 2-2m0 4v3h6v-3H5m8 0v3h6v-3h-6m-8 5v3h6v-3H5m8 0v3h6v-3h-6M11 2h10v4h-2V4h-6v2h-2V2Z"/>`),
		g.Group(children),
	)
}