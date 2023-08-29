package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineStartSquareOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 15.5h7v-7H4v7Zm-1 2q-.425 0-.713-.288T2 16.5v-9q0-.425.288-.713T3 6.5h9q.425 0 .713.288T13 7.5V11h8q.425 0 .713.288T22 12q0 .425-.288.713T21 13h-8v3.5q0 .425-.288.713T12 17.5H3ZM7.5 12Z"/>`),
		g.Group(children),
	)
}