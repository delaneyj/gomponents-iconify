package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickReferenceAllOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 4v6.025V10v10V4v5v-5Zm3 10h2.5q.225-.575.55-1.075T11.8 12H8q-.425 0-.713.288T7 13q0 .425.288.713T8 14Zm0 4h2.175q-.125-.5-.163-1t.013-1H8q-.425 0-.713.288T7 17q0 .425.288.713T8 18Zm-3 4q-.825 0-1.413-.588T3 20V4q0-.825.588-1.413T5 2h7.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V10.5q-.475-.2-.975-.313T17 10.025V9h-4q-.425 0-.713-.288T12 8V4H5v16h6.025q.4.6.9 1.113t1.1.887H5Zm11.5-3q1.05 0 1.775-.725T19 16.5q0-1.05-.725-1.775T16.5 14q-1.05 0-1.775.725T14 16.5q0 1.05.725 1.775T16.5 19Zm5.8 3.3q-.275.275-.7.275t-.7-.275l-2-2q-.525.35-1.137.525T16.5 21q-1.875 0-3.187-1.313T12 16.5q0-1.875 1.313-3.188T16.5 12q1.875 0 3.188 1.313T21 16.5q0 .65-.175 1.263T20.3 18.9l2 2q.275.275.275.7t-.275.7Z"/>`),
		g.Group(children),
	)
}