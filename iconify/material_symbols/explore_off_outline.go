package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExploreOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.9 12.05L11.95 9.1l5.55-2.6l-2.6 5.55Zm4.875 10.575L17.5 20.35q-1.225.8-2.613 1.225T12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-1.5.425-2.888T3.65 6.5L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425Zm.575-5.125l-1.45-1.45q.55-.95.825-1.963T20 12q0-3.325-2.338-5.663T12 4q-1.075 0-2.087.275T7.95 5.1L6.5 3.65q1.225-.8 2.613-1.225T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 1.5-.425 2.888T20.35 17.5Zm-6.925-6.925Zm-2.85 2.85ZM12 20q1.075 0 2.088-.275t1.962-.825l-4-4l-5.55 2.6l2.55-5.6L5.1 7.95q-.55.95-.825 1.962T4 12q0 3.325 2.337 5.663T12 20Z"/>`),
		g.Group(children),
	)
}