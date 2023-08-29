package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatAlignRightRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5q-.425 0-.713-.288T3 4q0-.425.288-.713T4 3h16q.425 0 .713.288T21 4q0 .425-.288.713T20 5H4Zm6 4q-.425 0-.713-.288T9 8q0-.425.288-.713T10 7h10q.425 0 .713.288T21 8q0 .425-.288.713T20 9H10Zm-6 4q-.425 0-.713-.288T3 12q0-.425.288-.713T4 11h16q.425 0 .713.288T21 12q0 .425-.288.713T20 13H4Zm6 4q-.425 0-.713-.288T9 16q0-.425.288-.713T10 15h10q.425 0 .713.288T21 16q0 .425-.288.713T20 17H10Zm-6 4q-.425 0-.713-.288T3 20q0-.425.288-.713T4 19h16q.425 0 .713.288T21 20q0 .425-.288.713T20 21H4Z"/>`),
		g.Group(children),
	)
}