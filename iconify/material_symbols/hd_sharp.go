package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 15h1.5v-2h2v2H11V9H9.5v2.5h-2V9H6v6Zm7 0h4.25l.75-.75v-4.5L17.25 9H13v6Zm1.5-1.5v-3h2v3h-2ZM2 20V4h20v16H2Z"/>`),
		g.Group(children),
	)
}