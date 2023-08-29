package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTwoWayLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.763 2.989L14.203.385c-.294-.3-.918-.3-1.211 0v1.814c-1.459.062-3.101.913-3.992 2.144c-.893-1.23-2.651-2.082-4.109-2.144V.385c-.295-.3-.918-.3-1.212 0L1.237 2.989a.776.776 0 0 0 0 1.085l2.442 2.603c.294.3.917.3 1.212 0V4.915c1.401.088 3.156 1.255 3.156 2.689v7.333h1.875V7.604c0-1.435 1.668-2.602 3.07-2.689v1.762c.293.3.917.3 1.211 0l2.56-2.603a.776.776 0 0 0 0-1.085z"/>`),
		g.Group(children),
	)
}