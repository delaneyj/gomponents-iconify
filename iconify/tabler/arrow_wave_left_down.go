package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowWaveLeftDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 14H3v-4"/><path d="M21 12c-.887 1.284-2.48 2.033-4 2c-1.52.033-3.113-.716-4-2s-2.48-2.033-4-2c-1.52-.033-3 1-4 2l-2 2"/></g>`),
		g.Group(children),
	)
}