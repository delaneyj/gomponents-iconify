package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 12V1H1v18h18v-7z"/><path fill="currentColor" d="M11 1v8h8V1zm6 6h-4V3h4z"/>`),
		g.Group(children),
	)
}