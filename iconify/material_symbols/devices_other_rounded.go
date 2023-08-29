package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesOtherRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h15q.425 0 .713.288T20 5q0 .425-.288.713T19 6H4v12h2q.425 0 .713.288T7 19q0 .425-.288.713T6 20H4Zm7-2.5q.625 0 1.063-.438T12.5 16q0-.625-.438-1.063T11 14.5q-.625 0-1.063.438T9.5 16q0 .625.438 1.063T11 17.5ZM21 20h-5q-.425 0-.713-.288T15 19v-9q0-.425.288-.713T16 9h5q.425 0 .713.288T22 10v9q0 .425-.288.713T21 20ZM9 19v-.775q-.475-.425-.738-1T8 16q0-.65.263-1.225t.737-1V13q0-.425.288-.713T10 12h2q.425 0 .713.288T13 13v.775q.475.425.738 1T14 16q0 .65-.263 1.225t-.737 1V19q0 .425-.288.713T12 20h-2q-.425 0-.713-.288T9 19Z"/>`),
		g.Group(children),
	)
}