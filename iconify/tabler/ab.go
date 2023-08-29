package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 16v-5.5a2.5 2.5 0 0 1 5 0V16m0-4H3m9-6v12m4-2V8h3a2 2 0 0 1 0 4h-3m3 0a2 2 0 0 1 0 4h-3"/>`),
		g.Group(children),
	)
}