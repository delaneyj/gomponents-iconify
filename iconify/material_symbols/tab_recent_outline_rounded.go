package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabRecentOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 23q-2.075 0-3.538-1.463T13 18q0-2.075 1.463-3.538T18 13q2.075 0 3.538 1.463T23 18q0 2.075-1.463 3.538T18 23Zm.5-5.2v-2.3q0-.2-.15-.35T18 15q-.2 0-.35.15t-.15.35v2.275q0 .2.075.388t.225.337l1.5 1.5q.15.15.35.15T20 20q.15-.15.15-.35T20 19.3l-1.5-1.5ZM4 20h7.3q-.15-.475-.225-.975T11 18H4V6h9v3q0 .425.288.713T14 10h6v1.3q.55.175 1.05.413t.95.562V6q0-.825-.588-1.413T20 4H4q-.825 0-1.413.588T2 6v12q0 .825.588 1.413T4 20Zm0-2V6v12Z"/>`),
		g.Group(children),
	)
}