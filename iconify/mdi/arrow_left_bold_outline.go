package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftBoldOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 22L3 12L13 2v6h8v8h-8v6M6 12l5 5v-3h8v-4h-8V7l-5 5Z"/>`),
		g.Group(children),
	)
}