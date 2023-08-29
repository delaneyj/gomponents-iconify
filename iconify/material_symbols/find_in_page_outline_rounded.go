package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FindInPageOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.75 20l2 2H6q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h8.075q.45 0 .85.188t.675.537l3.925 4.725q.225.275.35.6t.125.675V20q0 .5-.212.913t-.588.687L14 16.45q-.425.275-.925.413T12 17q-1.65 0-2.825-1.175T8 13q0-1.65 1.175-2.825T12 9q1.65 0 2.825 1.175T16 13q0 .575-.138 1.075T15.45 15L18 17.6V8.7L14.05 4H6v16h8.75ZM12 15q.825 0 1.413-.588T14 13q0-.825-.588-1.413T12 11q-.825 0-1.413.588T10 13q0 .825.588 1.413T12 15Zm0-2Zm0 0Z"/>`),
		g.Group(children),
	)
}