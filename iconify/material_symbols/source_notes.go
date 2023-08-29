package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SourceNotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 9h10V7H7v2Zm11 14q-1.825 0-3.188-1.137T13.1 19h1.55q.325 1.1 1.238 1.8t2.112.7q1.45 0 2.475-1.025T21.5 18q0-1.45-1.025-2.475T18 14.5q-.725 0-1.35.263t-1.1.737H17V17h-4v-4h1.5v1.425q.675-.65 1.575-1.038T18 13q2.075 0 3.538 1.463T23 18q0 2.075-1.463 3.538T18 23Zm-7.425-2H5q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v5.575q-.725-.275-1.475-.425T18 10q-1.05 0-2.025.262t-1.85.738H7v2h4.75q-.35.45-.65.95T10.575 15H7v2h3.05q-.05.25-.05.488V18q0 .775.15 1.525T10.575 21Z"/>`),
		g.Group(children),
	)
}