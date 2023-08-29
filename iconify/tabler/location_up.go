package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 18l-2-4l-7-3.5a.55.55 0 0 1 0-1L21 3l-3.251 9.003M19 22v-6m3 3l-3-3l-3 3"/>`),
		g.Group(children),
	)
}