package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatImageRightRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17q-.425 0-.713-.288T11 16V8q0-.425.288-.713T12 7h8q.425 0 .713.288T21 8v8q0 .425-.288.713T20 17h-8Zm1-2h6V9h-6v6Zm-9 6q-.425 0-.713-.288T3 20q0-.425.288-.713T4 19h16q.425 0 .713.288T21 20q0 .425-.288.713T20 21H4Zm0-4q-.425 0-.713-.288T3 16q0-.425.288-.713T4 15h4q.425 0 .713.288T9 16q0 .425-.288.713T8 17H4Zm0-4q-.425 0-.713-.288T3 12q0-.425.288-.713T4 11h4q.425 0 .713.288T9 12q0 .425-.288.713T8 13H4Zm0-4q-.425 0-.713-.288T3 8q0-.425.288-.713T4 7h4q.425 0 .713.288T9 8q0 .425-.288.713T8 9H4Zm0-4q-.425 0-.713-.288T3 4q0-.425.288-.713T4 3h16q.425 0 .713.288T21 4q0 .425-.288.713T20 5H4Z"/>`),
		g.Group(children),
	)
}