package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatIndentIncreaseRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.425 0-.713-.288T3 20q0-.425.288-.713T4 19h16q.425 0 .713.288T21 20q0 .425-.288.713T20 21H4Zm8-4q-.425 0-.713-.288T11 16q0-.425.288-.713T12 15h8q.425 0 .713.288T21 16q0 .425-.288.713T20 17h-8Zm0-4q-.425 0-.713-.288T11 12q0-.425.288-.713T12 11h8q.425 0 .713.288T21 12q0 .425-.288.713T20 13h-8Zm0-4q-.425 0-.713-.288T11 8q0-.425.288-.713T12 7h8q.425 0 .713.288T21 8q0 .425-.288.713T20 9h-8ZM4 5q-.425 0-.713-.288T3 4q0-.425.288-.713T4 3h16q.425 0 .713.288T21 4q0 .425-.288.713T20 5H4Zm-.15 10.15q-.25.25-.55.125T3 14.8V9.2q0-.35.3-.475t.55.125l2.8 2.8q.15.15.15.35t-.15.35l-2.8 2.8Z"/>`),
		g.Group(children),
	)
}