package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MathYMinusY(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 9l3 5.063M8 9l-4.8 9M16 9l3 5.063M22 9l-4.8 9M10 12h4"/>`),
		g.Group(children),
	)
}