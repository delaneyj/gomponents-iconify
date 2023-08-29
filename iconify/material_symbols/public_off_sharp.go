package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PublicOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.475 23.3l-2.95-2.95q-1.2.8-2.587 1.225T12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-1.55.425-2.938T3.65 6.476L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM11 19.95V18l-2-2v-1l-4.8-4.8q-.075.45-.138.9T4 12q0 3.025 1.988 5.3T11 19.95Zm9.35-2.475l-1.45-1.45q.525-.925.813-1.937T20 12q0-2.45-1.363-4.475T15 4.6V5l-4 2v1.125L6.525 3.65q1.2-.775 2.575-1.212T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 1.525-.438 2.9t-1.212 2.575Z"/>`),
		g.Group(children),
	)
}