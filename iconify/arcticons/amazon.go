package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amazon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M34.18 23.914c1.367-.554 3.8-1.29 4.533-.402c.792.96-.21 3.045-1.131 4.665"/><path d="M9 24.556c2.162 1.717 8.548 4.345 15.351 4.345c6.486 0 10.446-2.35 12.5-3.785"/></g>`),
		g.Group(children),
	)
}