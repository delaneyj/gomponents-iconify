package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DomainAddRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19V5q0-.825.588-1.413T4 3h6q.825 0 1.413.588T12 5v2h8q.825 0 1.413.588T22 9v5q0 .425-.288.713T21 15q-.425 0-.713-.288T20 14V9h-8v2h2v2h-2v2h2v2h-2v2h3q.425 0 .713.288T16 20q0 .425-.288.713T15 21H4Zm0-2h2v-2H4v2Zm0-4h2v-2H4v2Zm0-4h2V9H4v2Zm0-4h2V5H4v2Zm4 12h2v-2H8v2Zm0-4h2v-2H8v2Zm0-4h2V9H8v2Zm0-4h2V5H8v2Zm13 16q-.425 0-.713-.288T20 22v-1h-1q-.425 0-.713-.288T18 20q0-.425.288-.713T19 19h1v-1q0-.425.288-.713T21 17q.425 0 .713.288T22 18v1h1q.425 0 .713.288T24 20q0 .425-.288.713T23 21h-1v1q0 .425-.288.713T21 23Zm-5-10v-2h2v2h-2Zm0 4v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}