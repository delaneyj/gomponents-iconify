package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftRightBoldOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 16v6l10-10L14 2v6h-4V2L0 12l10 10v-6h4m-6-2v3l-5-5l5-5v3h8V7l5 5l-5 5v-3H8Z"/>`),
		g.Group(children),
	)
}