package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownBoldOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 11L12 21L2 11h6V3h8v8h6m-10 7l5-5h-3V5h-4v8H7l5 5Z"/>`),
		g.Group(children),
	)
}