package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Imgurviewer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.455 41.703L34.944 6.588h-8.18L22.02 22.464V17.14h-8.18v25.11a1.25 1.25 0 0 0 1.25 1.25h6.952a2.518 2.518 0 0 0 2.413-1.797Z"/><circle cx="17.99" cy="9.434" r="4.934" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}