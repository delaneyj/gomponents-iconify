package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DnaOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 2c-1.35 1.5-2.092 3-2.5 4.5M9 22c1.35-1.5 2.092-3 2.5-4.5M2 15c3.333-3 6.667-3 10-3m10-3c-1.5 1.35-3 2.092-4.5 2.5M17 6l-2.5-2.5M14 8l-1.5-1.5M7 18l2.5 2.5m-6-6l.5.5m16-6l.5.5m-14 3l1 1m9-3l1 1M10 16l1.5 1.5M2 2l20 20"/>`),
		g.Group(children),
	)
}