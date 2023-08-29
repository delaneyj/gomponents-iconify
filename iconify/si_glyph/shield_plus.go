package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 1.961v6.074c0 5.016 6.546 7.913 6.546 7.913s6.423-2.73 6.423-7.929c0-5.196.002-6.059.002-6.059S12.094.011 8.546.011S2 1.961 2 1.961zm9.031 6.101H9V10H7.938V8.062H6V7h1.938V4.938H9V7h2.031v1.062z"/>`),
		g.Group(children),
	)
}