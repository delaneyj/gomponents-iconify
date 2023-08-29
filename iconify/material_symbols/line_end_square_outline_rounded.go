package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineEndSquareOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 15.5h7v-7h-7v7Zm-1 2q-.425 0-.713-.288T11 16.5V13H3q-.425 0-.713-.288T2 12q0-.425.288-.713T3 11h8V7.5q0-.425.288-.713T12 6.5h9q.425 0 .713.288T22 7.5v9q0 .425-.288.713T21 17.5h-9Zm4.5-5.5Z"/>`),
		g.Group(children),
	)
}