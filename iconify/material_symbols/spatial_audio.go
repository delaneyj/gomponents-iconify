package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpatialAudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 10.025q-1.8 0-3.45-.688t-2.925-1.962Q15.35 6.1 14.662 4.45T13.976 1h2q0 1.425.525 2.7t1.525 2.275q1 1 2.275 1.538t2.7.537v1.975Zm0-3.975q-1.025 0-1.938-.375t-1.637-1.1q-.725-.725-1.1-1.638T17.95 1h1.975q0 .625.238 1.188t.662.987q.425.425.988.65T23 4.05v2ZM10 13q-1.65 0-2.825-1.175T6 9q0-1.65 1.175-2.825T10 5q1.65 0 2.825 1.175T14 9q0 1.65-1.175 2.825T10 13Zm-8 8v-2.8q0-.825.425-1.55t1.175-1.1q1.275-.65 2.875-1.1T10 14q1.925 0 3.525.45t2.875 1.1q.75.375 1.175 1.1T18 18.2V21H2Z"/>`),
		g.Group(children),
	)
}