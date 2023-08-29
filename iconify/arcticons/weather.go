package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 13a11 11 0 1 1-11 11a11 11 0 0 1 11-11Zm0-9.5v6.11m14.5-.11l-4.32 4.32M44.5 24h-6.11m.11 14.5l-4.32-4.32M24 44.5v-6.11m-14.5.11l4.32-4.32M3.5 24h6.11M9.5 9.5l4.32 4.32"/>`),
		g.Group(children),
	)
}