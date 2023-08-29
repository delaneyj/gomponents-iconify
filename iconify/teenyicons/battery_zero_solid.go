package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryZeroSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.5 13A1.5 1.5 0 0 1 0 11.5v-8A1.5 1.5 0 0 1 1.5 2h10A1.5 1.5 0 0 1 13 3.5v8a1.5 1.5 0 0 1-1.5 1.5h-10ZM15 10V5h-1v5h1Z"/>`),
		g.Group(children),
	)
}