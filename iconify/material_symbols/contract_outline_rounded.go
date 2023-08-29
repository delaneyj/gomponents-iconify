package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContractOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-1.25 0-2.125-.875T3 19v-1q0-.825.588-1.413T5 16h1V4q0-.825.588-1.413T8 2h11q.825 0 1.413.588T21 4v15q0 1.25-.875 2.125T18 22H6Zm12-2q.425 0 .713-.288T19 19V4H8v12h7q.825 0 1.413.588T17 18v1q0 .425.288.713T18 20ZM10 9q-.425 0-.713-.288T9 8q0-.425.288-.713T10 7h7q.425 0 .713.288T18 8q0 .425-.288.713T17 9h-7Zm0 3q-.425 0-.713-.288T9 11q0-.425.288-.713T10 10h7q.425 0 .713.288T18 11q0 .425-.288.713T17 12h-7Zm-4 8h9v-2H5v1q0 .425.288.713T6 20Zm0 0H5h10h-9Z"/>`),
		g.Group(children),
	)
}