package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscPlayTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.516 2.984L12.141 7h2.805L8.492.027L2.035 7h2.807l3.674-4.016zm0 10.035L4.809 9.031H2.003l6.468 6.903l6.468-6.903h-2.806l-3.617 3.988z"/>`),
		g.Group(children),
	)
}