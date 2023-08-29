package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NatureEcologyPottedPlantTreePlantSucculentPot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m11.5 7.5l-.86 5.16a1 1 0 0 1-1 .84H4.35a1 1 0 0 1-1-.84L2.5 7.5ZM4 .69a3.84 3.84 0 0 0-1.5 3.06A3.63 3.63 0 0 0 6 7.5a3.24 3.24 0 0 0 .94-.14Zm6 0a3.84 3.84 0 0 1 1.5 3.06A3.63 3.63 0 0 1 8 7.5a3.24 3.24 0 0 1-.94-.14Z"/><path d="M5 3L7 .5L9 3"/></g>`),
		g.Group(children),
	)
}