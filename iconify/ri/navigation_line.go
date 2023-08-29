package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.965 5.096l3.546 12.41l3.04-6.08l5.637-2.255L4.965 5.096ZM2.899 2.3l18.806 6.268a.5.5 0 0 1 .028.939L13 13l-4.425 8.85a.5.5 0 0 1-.928-.086L2.261 2.91a.5.5 0 0 1 .638-.611Z"/>`),
		g.Group(children),
	)
}