package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeedingOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.412 7.407a6.025 6.025 0 0 0-2.82-2.82M4 4H3v2a6 6 0 0 0 6 6h3m0 2a6 6 0 0 1 .255-1.736m1.51-2.514A5.981 5.981 0 0 1 18 8h3v1c0 2.158-1.14 4.05-2.85 5.107M15 15h-3m0 5v-8M3 3l18 18"/>`),
		g.Group(children),
	)
}