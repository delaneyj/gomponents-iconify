package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 6h12v12H6V6Zm12 0h12v12H18V6Zm0 12h12v12H18V18Zm0 12h12v12H18V30ZM30 6h12v12H30V6Z"/>`),
		g.Group(children),
	)
}