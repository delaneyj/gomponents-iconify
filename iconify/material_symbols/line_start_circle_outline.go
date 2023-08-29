package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineStartCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 16q1.65 0 2.825-1.175T12 12q0-1.65-1.175-2.825T8 8Q6.35 8 5.175 9.175T4 12q0 1.65 1.175 2.825T8 16Zm0 2q-2.5 0-4.25-1.75T2 12q0-2.5 1.75-4.25T8 6q2.25 0 3.913 1.425T13.925 11H22v2h-8.075q-.35 2.15-2.013 3.575T8 18Zm0-6Z"/>`),
		g.Group(children),
	)
}