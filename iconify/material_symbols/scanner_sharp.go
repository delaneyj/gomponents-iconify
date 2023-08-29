package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScannerSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20v-8h14.6L3.5 6.9L4.2 5L21 11.15V20H3Zm7-3h8v-2h-8v2Zm-4 0h2v-2H6v2Z"/>`),
		g.Group(children),
	)
}