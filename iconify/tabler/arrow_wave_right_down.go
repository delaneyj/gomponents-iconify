package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowWaveRightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 14h4v-4"/><path d="M3 12c.887 1.284 2.48 2.033 4 2c1.52.033 3.113-.716 4-2s2.48-2.033 4-2c1.52-.033 3 1 4 2l2 2"/></g>`),
		g.Group(children),
	)
}