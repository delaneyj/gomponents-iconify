package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatAlignCenterRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.425 0-.713-.288T3 20q0-.425.288-.713T4 19h16q.425 0 .713.288T21 20q0 .425-.288.713T20 21H4Zm4-4q-.425 0-.713-.288T7 16q0-.425.288-.713T8 15h8q.425 0 .713.288T17 16q0 .425-.288.713T16 17H8Zm-4-4q-.425 0-.713-.288T3 12q0-.425.288-.713T4 11h16q.425 0 .713.288T21 12q0 .425-.288.713T20 13H4Zm4-4q-.425 0-.713-.288T7 8q0-.425.288-.713T8 7h8q.425 0 .713.288T17 8q0 .425-.288.713T16 9H8ZM4 5q-.425 0-.713-.288T3 4q0-.425.288-.713T4 3h16q.425 0 .713.288T21 4q0 .425-.288.713T20 5H4Z"/>`),
		g.Group(children),
	)
}