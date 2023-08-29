package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineageClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="3.332" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 38.614V27.332m1.003-6.51l11.414-7.807a.288.288 0 0 1 .37.439l-9.545 9.78m-6.42-.236l-6.27-7.438a.29.29 0 0 1 .394-.422l7.826 5.765"/>`),
		g.Group(children),
	)
}