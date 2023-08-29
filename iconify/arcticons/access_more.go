package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessMore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 18.011l7.041 7.04L24 32.092l-7.041-7.041L24 18.011"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.898 9.452l15.599 15.599L24 38.548L10.503 25.051l7.034-7.033"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.898 3l22.051 22.051L24 45L4.051 25.051l7.034-7.033"/>`),
		g.Group(children),
	)
}