package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageFilter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 5.87c0-2.2-1.79-4-4-4s-4 1.8-4 4c0 2.21 1.79 4 4 4s4-1.79 4-4zM3.24 10.66a3.995 3.995 0 0 0-1.47 5.46a4.006 4.006 0 0 0 5.47 1.47a4.003 4.003 0 0 0 1.46-5.47c-1.1-1.91-3.55-2.56-5.46-1.46zm9.52 6.93c1.92 1.1 4.36.45 5.47-1.46c1.1-1.92.45-4.36-1.47-5.47c-1.91-1.1-4.36-.45-5.46 1.46a4.003 4.003 0 0 0 1.46 5.47z"/>`),
		g.Group(children),
	)
}