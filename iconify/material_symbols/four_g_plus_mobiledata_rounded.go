package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourGPlusMobiledataRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 13h-1q-.425 0-.713-.288T18 12q0-.425.288-.713T19 11h1v-1q0-.425.288-.713T21 9q.425 0 .713.288T22 10v1h1q.425 0 .713.288T24 12q0 .425-.288.713T23 13h-1v1q0 .425-.288.713T21 15q-.425 0-.713-.288T20 14v-1Zm-9 4q-.825 0-1.413-.588T9 15V9q0-.825.588-1.413T11 7h5q.425 0 .713.288T17 8q0 .425-.288.713T16 9h-5v6h4v-2h-1q-.425 0-.713-.288T13 12q0-.425.288-.713T14 11h2q.425 0 .713.288T17 12v3q0 .825-.588 1.413T15 17h-4ZM3 8v4h2V8q0-.425.288-.713T6 7q.425 0 .713.288T7 8v4q.425 0 .713.288T8 13q0 .425-.288.713T7 14v2q0 .425-.288.713T6 17q-.425 0-.713-.288T5 16v-2H2q-.425 0-.713-.288T1 13V8q0-.425.288-.713T2 7q.425 0 .713.288T3 8Z"/>`),
		g.Group(children),
	)
}