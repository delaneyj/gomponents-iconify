package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leaf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M37 24c0 14.094-13 20-13 20s-13-4.625-13-20S24 4 24 4s13 5.906 13 20ZM24 36l5-5m-5-2l-5-5m5-1l5-5m-5 26V14"/>`),
		g.Group(children),
	)
}