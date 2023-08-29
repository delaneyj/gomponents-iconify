package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upstair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.969 1.047h-2.896v2h-2.011v2h-2.01v2H8.043v2H6.031v2h-2.01v2H1v2.918h2.896L17 3.965l-.031-2.918zm-9.096.057H2.949l2.155 1.414l-3.182 3.18l1.326 1.326l3.181-3.181l1.448 2.205l-.004-4.944z"/>`),
		g.Group(children),
	)
}