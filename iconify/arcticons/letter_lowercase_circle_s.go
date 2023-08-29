package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterLowercaseCircleS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.427 30.125c1.015.852 2.111 1.243 4.573 1.243h1.248a3.68 3.68 0 0 0 3.675-3.684h0A3.68 3.68 0 0 0 25.247 24h-2.495a3.68 3.68 0 0 1-3.675-3.684h0a3.68 3.68 0 0 1 3.675-3.684H24c2.462 0 3.558.39 4.573 1.243"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}