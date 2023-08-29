package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdfScannerSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-8h4V4h12v8h4v8H2Zm6-8h8V6H8v6Zm10 5q.425 0 .713-.288T19 16q0-.425-.288-.713T18 15q-.425 0-.713.288T17 16q0 .425.288.713T18 17Z"/>`),
		g.Group(children),
	)
}