package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSelectMoveDownRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5q-.425 0-.713-.288T3 4q0-.425.288-.713T4 3h16q.425 0 .713.288T21 4q0 .425-.288.713T20 5H4Zm7.3 11.3l-2.6-2.6q-.275-.275-.287-.687T8.7 12.3q.275-.275.7-.275t.7.275l.9.875V8q0-.425.288-.713T12 7q.425 0 .713.288T13 8v5.175l.9-.9q.275-.275.688-.275t.712.3q.275.275.275.7t-.275.7l-2.6 2.6q-.3.3-.7.3t-.7-.3ZM4 21q-.425 0-.713-.288T3 20q0-.425.288-.713T4 19h16q.425 0 .713.288T21 20q0 .425-.288.713T20 21H4Z"/>`),
		g.Group(children),
	)
}