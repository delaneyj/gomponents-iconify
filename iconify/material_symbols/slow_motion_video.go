package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlowMotionVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.25 18.3q-.95-1.1-1.525-2.45T2 13h2.05q.15 1.1.55 2.087T5.65 16.9l-1.4 1.4ZM2 11q.2-1.5.75-2.85t1.5-2.45l1.4 1.4Q5 7.925 4.6 8.913T4.05 11H2Zm8.95 10.95q-1.5-.15-2.837-.725T5.65 19.75l1.4-1.45q.875.65 1.838 1.075t2.062.575v2ZM7.1 5.7L5.65 4.25q1.125-.9 2.463-1.475T11 2.05v2q-1.125.15-2.1.575T7.1 5.7Zm2.4 10.8v-9l7 4.5l-7 4.5Zm3.5 5.45v-2q3.025-.425 5.013-2.675T20 12q0-3.025-1.988-5.275T13 4.05v-2q3.85.425 6.425 3.25T22 12q0 3.875-2.575 6.7T13 21.95Z"/>`),
		g.Group(children),
	)
}