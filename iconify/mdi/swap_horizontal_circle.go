package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapHorizontalCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 12c0-5.5-4.5-10-10-10S2 6.5 2 12s4.5 10 10 10s10-4.5 10-10m-7-5.5l3.5 3.5l-3.5 3.5V11h-4V9h4V6.5m-6 11L5.5 14L9 10.5V13h4v2H9v2.5Z"/>`),
		g.Group(children),
	)
}