package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Molotov(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 24A21.5 21.5 0 1 1 24 2.5A21.51 21.51 0 0 1 45.5 24Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 18.18c1.59-3.2 4.85-6.7 5.6-6.7s-1.64 9.15-2.8 21c0 0 5.91-17.82 9.53-17.82c1.44 0-1.45 13.88-1.45 13.88s5.9-14.08 9.49-14.08c2.39 0-2.17 13.56-2.09 16.69a11.7 11.7 0 0 0 1.5 5.41m1.28-3.21a24.49 24.49 0 0 1 5.94 1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.88 36a79.07 79.07 0 0 1 13.56-2.52"/>`),
		g.Group(children),
	)
}