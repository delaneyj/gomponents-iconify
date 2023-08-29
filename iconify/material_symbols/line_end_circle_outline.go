package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineEndCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 16q1.65 0 2.825-1.175T20 12q0-1.65-1.175-2.825T16 8q-1.65 0-2.825 1.175T12 12q0 1.65 1.175 2.825T16 16Zm0 2q-2.25 0-3.913-1.425T10.075 13H2v-2h8.075q.35-2.15 2.013-3.575T16 6q2.5 0 4.25 1.75T22 12q0 2.5-1.75 4.25T16 18Zm0-6Z"/>`),
		g.Group(children),
	)
}