package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinInvokeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h9q.425 0 .713.288T14 5q0 .425-.288.713T13 6H4v12h16v-5q0-.425.288-.713T21 12q.425 0 .713.288T22 13v5q0 .825-.588 1.413T20 20H4Zm9-6.575l-2.25 2.25q-.3.3-.7.288t-.7-.313q-.275-.3-.288-.7t.288-.7L11.6 12h-1.25q-.425 0-.713-.287T9.35 11q0-.425.288-.713T10.35 10H14q.425 0 .713.288T15 11v3.65q0 .425-.288.713T14 15.65q-.425 0-.713-.288T13 14.65v-1.225ZM19 10q-1.25 0-2.125-.875T16 7q0-1.25.875-2.125T19 4q1.25 0 2.125.875T22 7q0 1.25-.875 2.125T19 10Z"/>`),
		g.Group(children),
	)
}