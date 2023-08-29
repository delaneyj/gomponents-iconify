package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineStartCircleRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 18q-2.5 0-4.25-1.75T2 12q0-2.5 1.75-4.25T8 6q2.25 0 3.913 1.425T13.925 11H21q.425 0 .713.288T22 12q0 .425-.288.713T21 13h-7.075q-.35 2.15-2.013 3.575T8 18Z"/>`),
		g.Group(children),
	)
}