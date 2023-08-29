package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MetroTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.5 42.38h-11a2 2 0 0 1-2-2V7.62a2 2 0 0 1 2-2h2.37a3 3 0 0 1 2.21 1l5.36 5.81a4 4 0 0 1 1.06 2.67Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.87 5.62h13a3 3 0 0 1 2.21 1l5.36 5.81a4 4 0 0 1 1.06 2.67v27.28h-13"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.87 5.62h13a3 3 0 0 1 2.21 1l5.36 5.81a4 4 0 0 1 1.06 2.67v25.28a2 2 0 0 1-2 2h-11"/>`),
		g.Group(children),
	)
}