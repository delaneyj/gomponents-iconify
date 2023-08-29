package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LtePlusMobiledataRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 16H2q-.425 0-.713-.288T1 15V9q0-.425.288-.713T2 8q.425 0 .713.288T3 9v5h2q.425 0 .713.288T6 15q0 .425-.288.713T5 16Zm3 0q-.425 0-.713-.288T7 15v-5H6q-.425 0-.713-.288T5 9q0-.425.288-.713T6 8h4q.425 0 .713.288T11 9q0 .425-.288.713T10 10H9v5q0 .425-.288.713T8 16Zm8 0h-3q-.425 0-.713-.288T12 15V9q0-.425.288-.713T13 8h3q.425 0 .713.288T17 9q0 .425-.288.713T16 10h-2v1h2q.425 0 .713.288T17 12q0 .425-.288.713T16 13h-2v1h2q.425 0 .713.288T17 15q0 .425-.288.713T16 16Zm4-3h-1q-.425 0-.713-.288T18 12q0-.425.288-.713T19 11h1v-1q0-.425.288-.713T21 9q.425 0 .713.288T22 10v1h1q.425 0 .713.288T24 12q0 .425-.288.713T23 13h-1v1q0 .425-.288.713T21 15q-.425 0-.713-.288T20 14v-1Z"/>`),
		g.Group(children),
	)
}