package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoNotDisturbOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.625L17.5 20.35q-1.225.8-2.612 1.225T12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-1.5.425-2.888T3.65 6.5L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425ZM7 13h3.15l-2-2H7v2Zm13.35 4.5l-4.5-4.5H17v-2h-3.15L6.5 3.65q1.225-.8 2.613-1.225T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 1.5-.425 2.888T20.35 17.5Z"/>`),
		g.Group(children),
	)
}