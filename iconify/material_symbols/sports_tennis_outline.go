package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsTennisOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.425 20L2 18.6l4.1-4.1q.775-.775 1.063-1.938T7.45 9q0-1.45.65-2.85t1.85-2.6q2.275-2.275 5.025-2.575T19.5 2.5q1.8 1.8 1.5 4.55t-2.55 5q-1.2 1.2-2.6 1.85t-2.85.65q-2.425 0-3.55.275T7.525 15.9l-4.1 4.1Zm6.875-8.35q1.175 1.15 3.175.85t3.575-1.875q1.6-1.6 1.913-3.588T18.1 3.925q-1.2-1.2-3.138-.9t-3.562 1.9Q9.825 6.5 9.487 8.488t.813 3.162ZM18 23q-1.65 0-2.825-1.175T14 19q0-1.65 1.175-2.825T18 15q1.65 0 2.825 1.175T22 19q0 1.65-1.175 2.825T18 23Zm0-2q.825 0 1.413-.588T20 19q0-.825-.588-1.413T18 17q-.825 0-1.413.588T16 19q0 .825.588 1.413T18 21Zm0-2Z"/>`),
		g.Group(children),
	)
}