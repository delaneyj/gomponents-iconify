package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoThousandFifty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.29 10.84a6.34 6.34 0 1 0-12.67 0v.28a14.06 14.06 0 0 0-16.38 1.41a2.9 2.9 0 1 0-3.26 4a14.1 14.1 0 0 0 12.45 20.79a13.76 13.76 0 0 0 4.13-.62a4.89 4.89 0 1 0 5.18-2.91a14 14 0 0 0 3.38-16.68a5.69 5.69 0 0 0 .83.07a6.34 6.34 0 0 0 6.34-6.34Z"/>`),
		g.Group(children),
	)
}