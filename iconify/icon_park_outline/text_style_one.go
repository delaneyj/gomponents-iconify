package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextStyleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m4 42l4.941-12M32 42l-4.941-12m0 0L25 25L18 8l-7 17l-2.059 5m18.118 0H8.94M28 10h16M32 20h12m-8 10h8m-4 10h4"/>`),
		g.Group(children),
	)
}