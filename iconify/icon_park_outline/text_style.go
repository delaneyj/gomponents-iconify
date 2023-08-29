package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextStyle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M32 6h10v10M17 32l2.188-5M31 32l-2.188-5m-9.625 0L24 16l4.813 11m-9.625 0h9.625M16 6H6v10m26 26h10V32M16 42H6V32"/>`),
		g.Group(children),
	)
}