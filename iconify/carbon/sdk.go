package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sdk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 9h-2.1L24 15.6V9h-2v14h2v-4.3l.9-1.5l3 5.8H30l-3.9-7.6L30 9zM16 23h-4V9h4a4.012 4.012 0 0 1 4 4v6a4.012 4.012 0 0 1-4 4zm-2-2h2a2.006 2.006 0 0 0 2-2v-6a2.006 2.006 0 0 0-2-2h-2zm-6 2H2v-2h6v-4H4a2.006 2.006 0 0 1-2-2v-4a2.006 2.006 0 0 1 2-2h6v2H4v4h4a2.006 2.006 0 0 1 2 2v4a2.006 2.006 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}