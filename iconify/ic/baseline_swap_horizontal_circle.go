package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineSwapHorizontalCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 12c0-5.52-4.48-10-10-10S2 6.48 2 12s4.48 10 10 10s10-4.48 10-10zm-7-5.5l3.5 3.5l-3.5 3.5V11h-4V9h4V6.5zm-6 11L5.5 14L9 10.5V13h4v2H9v2.5z"/>`),
		g.Group(children),
	)
}