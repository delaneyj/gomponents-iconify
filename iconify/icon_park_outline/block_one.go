package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 30h12v12H6V30ZM30 6h12v12H30V6Zm0 12h12v12H30V18Zm-12 0h12v12H18V18ZM6 18h12v12H6V18Z"/>`),
		g.Group(children),
	)
}