package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.026 24a6.075 6.075 0 0 0-6.075 6.075h0a6.075 6.075 0 0 0 6.075 6.075h3.948a6.075 6.075 0 0 0 6.075-6.075h0A6.075 6.075 0 0 0 25.974 24m0 0a6.075 6.075 0 0 0 6.075-6.075h0a6.075 6.075 0 0 0-6.075-6.075h-3.948a6.075 6.075 0 0 0-6.075 6.075h0A6.075 6.075 0 0 0 22.026 24m0 0h3.948"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}