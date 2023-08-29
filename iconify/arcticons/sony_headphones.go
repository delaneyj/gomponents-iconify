package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SonyHeadphones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m45.028 28.484l-3.707-1.728a2 2 0 0 0-2.657.967l-3.17 6.798a2 2 0 0 0 .967 2.658l3.218 1.5C43.279 34.835 45.5 29.683 45.5 24c0-11.874-9.626-21.5-21.5-21.5S2.5 12.126 2.5 24c0 5.682 2.22 10.835 5.821 14.679l3.218-1.5a2 2 0 0 0 .967-2.658l-3.17-6.798a2 2 0 0 0-2.657-.967l-3.707 1.728"/>`),
		g.Group(children),
	)
}