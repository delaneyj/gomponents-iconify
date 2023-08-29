package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CisojrFourDroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.163 41.545A17.648 17.648 0 0 1 6.517 27.463M22.909 6.48q.466-.025.933-.025a17.648 17.648 0 0 1 17.642 17.198"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.909 9.69q.258-.01.516-.01A13.552 13.552 0 0 1 36.99 23.22h0q0 .217-.007.434M21.117 36.562a13.557 13.557 0 0 1-10.583-9.13"/><ellipse cx="23.841" cy="23.658" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.364" ry="2.503"/><circle cx="23.841" cy="23.652" r="6.019" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}