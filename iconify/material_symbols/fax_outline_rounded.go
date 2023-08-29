package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaxOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-1.25 0-2.125-.875T2 18v-7q0-1.25.875-2.125T5 8q.675 0 1.238.275T7.224 9H8V6q0-.825.588-1.413T10 4h6q.825 0 1.413.588T18 6v3h1q1.25 0 2.125.875T22 12v6q0 .825-.588 1.413T20 20H7.225q-.425.45-.987.725T5 21Zm0-2q.425 0 .713-.288T6 18v-7q0-.425-.288-.713T5 10q-.425 0-.713.288T4 11v7q0 .425.288.713T5 19Zm5-10h6V6h-6v3Zm-2 9h12v-6q0-.425-.288-.713T19 11H8v7Zm7-4q.425 0 .713-.288T16 13q0-.425-.288-.713T15 12q-.425 0-.713.288T14 13q0 .425.288.713T15 14Zm3 0q.425 0 .713-.288T19 13q0-.425-.288-.713T18 12q-.425 0-.713.288T17 13q0 .425.288.713T18 14Zm-3 3q.425 0 .713-.288T16 16q0-.425-.288-.713T15 15q-.425 0-.713.288T14 16q0 .425.288.713T15 17Zm3 0q.425 0 .713-.288T19 16q0-.425-.288-.713T18 15q-.425 0-.713.288T17 16q0 .425.288.713T18 17Zm-9 0h4v-5H9v5Zm-1 1v-7v7Z"/>`),
		g.Group(children),
	)
}