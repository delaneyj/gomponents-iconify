package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TokenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 22l-9-5V7l9-5l9 5v10l-9 5ZM9.1 9.25q.575-.6 1.325-.925T12 8q.825 0 1.575.325t1.325.925l3-1.675L12 4.3L6.1 7.575l3 1.675Zm1.9 9.9v-3.275q-1.35-.35-2.175-1.425T8 12q0-.275.025-.513t.1-.487L5 9.25v6.575l6 3.325ZM12 14q.825 0 1.413-.588T14 12q0-.825-.588-1.413T12 10q-.825 0-1.413.588T10 12q0 .825.588 1.413T12 14Zm1 5.15l6-3.325V9.25L15.875 11q.075.25.1.488T16 12q0 1.375-.825 2.45T13 15.875v3.275Z"/>`),
		g.Group(children),
	)
}