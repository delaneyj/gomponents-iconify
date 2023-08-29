package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChairOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21v-2H1V8h3V3h16v5h3v11h-3v2h-2v-2H6v2H4Zm-1-4h18v-7h-2v5H5v-5H3v7Zm4-4h10V8h1V5H6v3h1v5Zm5 2Zm0-2Zm0 2Z"/>`),
		g.Group(children),
	)
}