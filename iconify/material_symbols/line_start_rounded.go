package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineStartRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.5 14.5q-1.05 0-1.775-.725T2 12q0-1.05.725-1.775T4.5 9.5q.775 0 1.4.425T6.8 11H21q.425 0 .713.288T22 12q0 .425-.288.713T21 13H6.8q-.275.65-.9 1.075t-1.4.425Z"/>`),
		g.Group(children),
	)
}