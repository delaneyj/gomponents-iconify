package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdfScanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-5q0-1.25.875-2.125T5 12h1V4h12v8h1q1.25 0 2.125.875T22 15v5H2Zm6-8h8V6H8v6Zm10 5q.425 0 .713-.288T19 16q0-.425-.288-.713T18 15q-.425 0-.713.288T17 16q0 .425.288.713T18 17Z"/>`),
		g.Group(children),
	)
}