package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exodus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.96 5.5L24 21.46L8.04 5.5M42.5 39.96L26.54 24L42.5 8.04M8.04 42.5L24 26.54L39.96 42.5M5.5 8.04L21.46 24L5.5 39.96"/>`),
		g.Group(children),
	)
}