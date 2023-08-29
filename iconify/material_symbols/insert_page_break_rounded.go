package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertPageBreakRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20v-2q0-.425.288-.713T5 17h14q.425 0 .713.288T20 18v2q0 .825-.588 1.413T18 22H6Zm7-18v4q0 .425.288.713T14 9h4l-5-5ZM4 4q0-.825.588-1.413T6 2h7.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V10q0 .425-.288.713T19 11H5q-.425 0-.713-.288T4 10V4Zm6 11q-.425 0-.713-.288T9 14q0-.425.288-.713T10 13h4q.425 0 .713.288T15 14q0 .425-.288.713T14 15h-4Zm8 0q-.425 0-.713-.288T17 14q0-.425.288-.713T18 13h4q.425 0 .713.288T23 14q0 .425-.288.713T22 15h-4ZM2 15q-.425 0-.713-.288T1 14q0-.425.288-.713T2 13h4q.425 0 .713.288T7 14q0 .425-.288.713T6 15H2Z"/>`),
		g.Group(children),
	)
}