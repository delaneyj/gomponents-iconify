package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquircleIde(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.861 17.639L3.5 24L24 44.5l6.361-6.361m7.778-7.778L44.5 24L24 3.5l-6.361 6.361"/>`),
		g.Group(children),
	)
}