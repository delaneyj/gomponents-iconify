package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutosumTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.81 3.706a.75.75 0 0 1 .69-.456h11a.75.75 0 0 1 0 1.5H6.262l4.146 4.308a.75.75 0 0 1 .035 1.001L6.104 15.25H15.5a.75.75 0 0 1 0 1.5h-11a.75.75 0 0 1-.575-1.231L8.86 9.613L3.96 4.52a.75.75 0 0 1-.15-.814Z"/>`),
		g.Group(children),
	)
}