package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoNotDisturbOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.85 13l-2-2H17v2h-1.15Zm3.925 9.625L17.5 20.35q-1.225.8-2.612 1.225T12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-1.5.425-2.888T3.65 6.5L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425Zm.575-5.125l-1.45-1.45q.55-.95.825-1.963T20 12q0-3.325-2.337-5.663T12 4q-1.075 0-2.087.275T7.95 5.1L6.5 3.65q1.225-.8 2.613-1.225T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 1.5-.425 2.888T20.35 17.5Zm-4.3 1.4l-5.9-5.9H7v-2h1.15L5.1 7.95q-.55.95-.825 1.963T4 12q0 3.325 2.337 5.663T12 20q1.075 0 2.087-.275t1.963-.825Zm-2.2-7.9Zm-3.275 2.425Z"/>`),
		g.Group(children),
	)
}