package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataUsage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-3.875 2.575-6.725T11 2.05v3q-2.6.35-4.3 2.325T5 12q0 2.925 2.038 4.963T12 19q1.6 0 3.038-.675T17.5 16.4l2.6 1.5q-1.425 1.95-3.55 3.025T12 22Zm9.15-5.95l-2.6-1.5q.225-.625.337-1.263T19 12q0-2.65-1.7-4.625T13 5.05v-3q3.85.375 6.425 3.225T22 12q0 1.05-.188 2.075t-.662 1.975Z"/>`),
		g.Group(children),
	)
}