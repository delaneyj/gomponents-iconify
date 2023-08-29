package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AuthZero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.98 7.448L19.62 0H4.347L2.02 7.448c-1.352 4.312.03 9.206 3.815 12.015L12.007 24l6.157-4.552c3.755-2.81 5.182-7.688 3.815-12.015l-6.16 4.58l2.343 7.45l-6.157-4.597l-6.158 4.58l2.358-7.433l-6.188-4.55l7.63-.045L12.008 0l2.356 7.404l7.615.044z"/>`),
		g.Group(children),
	)
}