package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trophy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 5.062V3H6v2.062H2V8c0 2.525 1.889 4.598 4.324 4.932A5.998 5.998 0 0 0 11 16.91V18a2 2 0 0 1-2 2H8v2h8v-2h-1a2 2 0 0 1-2-2v-1.09a5.998 5.998 0 0 0 4.676-3.978C20.111 12.598 22 10.525 22 8V5.062h-4zM4 8v-.938h2v3.766A3.004 3.004 0 0 1 4 8zm16 0a3.006 3.006 0 0 1-2 2.829V7.062h2V8z"/>`),
		g.Group(children),
	)
}