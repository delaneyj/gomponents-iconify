package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineEnd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 14.5q-.775 0-1.4-.425T17.2 13H2v-2h15.2q.275-.65.9-1.075t1.4-.425q1.05 0 1.775.725T22 12q0 1.05-.725 1.775T19.5 14.5Z"/>`),
		g.Group(children),
	)
}