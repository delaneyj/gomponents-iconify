package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatImageLeftRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17q-.425 0-.713-.288T3 16V8q0-.425.288-.713T4 7h8q.425 0 .713.288T13 8v8q0 .425-.288.713T12 17H4Zm1-2h6V9H5v6ZM4 5q-.425 0-.713-.288T3 4q0-.425.288-.713T4 3h16q.425 0 .713.288T21 4q0 .425-.288.713T20 5H4Zm12 4q-.425 0-.713-.288T15 8q0-.425.288-.713T16 7h4q.425 0 .713.288T21 8q0 .425-.288.713T20 9h-4Zm0 4q-.425 0-.713-.288T15 12q0-.425.288-.713T16 11h4q.425 0 .713.288T21 12q0 .425-.288.713T20 13h-4Zm0 4q-.425 0-.713-.288T15 16q0-.425.288-.713T16 15h4q.425 0 .713.288T21 16q0 .425-.288.713T20 17h-4ZM4 21q-.425 0-.713-.288T3 20q0-.425.288-.713T4 19h16q.425 0 .713.288T21 20q0 .425-.288.713T20 21H4Z"/>`),
		g.Group(children),
	)
}