package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 14h6q.425 0 .713-.288T14 13q0-.425-.288-.713T13 12H7q-.425 0-.713.288T6 13q0 .425.288.713T7 14Zm0-3h10q.425 0 .713-.288T18 10q0-.425-.288-.713T17 9H7q-.425 0-.713.288T6 10q0 .425.288.713T7 11Zm0-3h10q.425 0 .713-.288T18 7q0-.425-.288-.713T17 6H7q-.425 0-.713.288T6 7q0 .425.288.713T7 8ZM6 18l-2.3 2.3q-.475.475-1.088.213T2 19.575V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H6Zm-2-2h16V4H4v12Zm0 0V4v12Z"/>`),
		g.Group(children),
	)
}