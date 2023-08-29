package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vrs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.37 22.615v-6.808h2.213a2.298 2.298 0 1 1 0 4.595h-2.212m2.375-.007l2.135 2.22m-6.254-6.808l-2.212 6.808l-2.298-6.808m12.556 6.042a2.029 2.029 0 0 0 1.702.766h1.021a1.702 1.702 0 0 0 0-3.404H23.29a1.702 1.702 0 1 1 0-3.404h1.021a1.827 1.827 0 0 1 1.702.766"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 24A21.5 21.5 0 1 1 24 2.5A21.5 21.5 0 0 1 45.5 24Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 6.781a17.22 17.22 0 0 0-13.383 6.436h14.736a5.95 5.95 0 0 1 0 11.898H6.859A17.199 17.199 0 1 0 24 6.781Z"/>`),
		g.Group(children),
	)
}