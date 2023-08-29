package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinimizeBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2 22l.875-.875M9 15H3.143M9 15v5.857M9 15l-3.5 3.5M22 2l-7 7m0 0h5.857M15 9V3.143"/>`),
		g.Group(children),
	)
}