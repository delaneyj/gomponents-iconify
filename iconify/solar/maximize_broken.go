package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaximizeBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 2h-5.857M22 2v5.857M22 2l-3.5 3.5M15 9l.875-.875M9 15l-7 7m0 0h5.857M2 22v-5.857"/>`),
		g.Group(children),
	)
}