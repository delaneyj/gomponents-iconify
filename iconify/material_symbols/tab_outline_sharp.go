package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18h16v-8h-7V6H4v12Zm-2 2V4h20v16H2Zm2-2V6v12Z"/>`),
		g.Group(children),
	)
}