package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pisces(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 24h40M10 7s6 9.82 6 17s-6 17-6 17M38 7s-6 9.96-6 17c0 7.04 6 17 6 17"/>`),
		g.Group(children),
	)
}