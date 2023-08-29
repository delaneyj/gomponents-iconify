package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewRows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3H2zm0 2v4h20v-4H2zm20 6H2a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}