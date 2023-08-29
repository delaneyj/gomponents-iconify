package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Materialistic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.45 8.4a2 2 0 0 0-1.95 1.95v27.3a2 2 0 0 0 2 2h35.1a2 2 0 0 0 2-2v-27.3a2 2 0 0 0-2-1.95Zm3.79 8.14l1.71-3l1.71 3m0 2.41l-1.71 3l-1.71-3m0 10.08l1.71-3l1.71 3m0 2.41l-1.71 3l-1.71-3m7.68-17.86h21.11m0 4.17H21.65m3.72 4.16h13.66m-21.11 4.18h21.11m0 4.16H21.65m3.72 4.17h13.66"/>`),
		g.Group(children),
	)
}