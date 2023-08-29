package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalFourGPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 12h4M3 8v3a1 1 0 0 0 1 1h3m0-4v8m12-6v4m-5-6h-2a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2v-4h-1"/>`),
		g.Group(children),
	)
}