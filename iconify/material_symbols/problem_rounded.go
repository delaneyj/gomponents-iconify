package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProblemRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 16q.425 0 .713-.288T8 15q0-.425-.288-.713T7 14q-.425 0-.713.288T6 15q0 .425.288.713T7 16Zm0-3q.425 0 .713-.288T8 12V9q0-.425-.288-.713T7 8q-.425 0-.713.288T6 9v3q0 .425.288.713T7 13Zm4 2h6q.425 0 .713-.288T18 14q0-.425-.288-.713T17 13h-6q-.425 0-.713.288T10 14q0 .425.288.713T11 15Zm0-4h6q.425 0 .713-.288T18 10q0-.425-.288-.713T17 9h-6q-.425 0-.713.288T10 10q0 .425.288.713T11 11Zm-7 9q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Z"/>`),
		g.Group(children),
	)
}