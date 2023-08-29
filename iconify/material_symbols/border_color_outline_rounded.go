package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderColorOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 24q-.825 0-1.413-.588T2 22q0-.825.588-1.413T4 20h16q.825 0 1.413.588T22 22q0 .825-.588 1.413T20 24H4Zm1-6q-.425 0-.713-.288T4 17v-2.325q0-.2.075-.388t.225-.337l8.75-8.75l3.75 3.75l-8.75 8.75q-.15.15-.338.225T7.325 18H5Zm1-2h.9L14 8.95L13.05 8L6 15.1v.9Zm11.925-8.15l-3.75-3.75l1.8-1.8q.275-.3.7-.287t.7.287l2.35 2.35q.275.275.275.688t-.275.712l-1.8 1.8ZM6 16Z"/>`),
		g.Group(children),
	)
}