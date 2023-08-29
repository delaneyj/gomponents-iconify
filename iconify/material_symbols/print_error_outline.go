package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintErrorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21v-4H2v-6q0-1.275.875-2.138T5 8h14q1 0 1.763.563T21.825 10H5q-.425 0-.713.288T4 11v4h2v-2h10v2H8v4h8v2H6ZM6 8V3h12v5h-2V5H8v3H6Zm13 13q-.425 0-.713-.288T18 20q0-.425.288-.713T19 19q.425 0 .713.288T20 20q0 .425-.288.713T19 21Zm-1-4v-5h2v5h-2ZM4 10h17.825H4Z"/>`),
		g.Group(children),
	)
}