package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M2.5 1A1.5 1.5 0 0 0 1 2.5v8A1.5 1.5 0 0 0 2.5 12h10a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 12.5 1h-10ZM0 14h15v-1H0v1Z"/>`),
		g.Group(children),
	)
}