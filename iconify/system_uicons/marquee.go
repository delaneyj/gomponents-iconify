package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Marquee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 5.5v-1a2 2 0 0 1 2-2h1m0 16h-1a2 2 0 0 1-2-2v-1m16-10v-1a2 2 0 0 0-2-2h-1m0 16h1a2 2 0 0 0 2-2v-1m-11-13h2m2 0h2m-6 16h2m2 0h2m5-11.002V9.5m0 1.998V13.5m-16-6.002V9.5m0 1.998V13.5"/>`),
		g.Group(children),
	)
}