package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsBath(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M21 10H7V7.113c0-.997.678-1.923 1.661-2.085A2.003 2.003 0 0 1 11 7h2a4.003 4.003 0 0 0-4.398-3.98C6.523 3.222 5 5.089 5 7.178V10H3a1 1 0 0 0-1 1v2c0 2.606 1.674 4.823 4 5.65V22h2v-3h8v3h2v-3.35c2.326-.827 4-3.044 4-5.65v-2a1 1 0 0 0-1-1z" fill="currentColor"/>`),
		g.Group(children),
	)
}