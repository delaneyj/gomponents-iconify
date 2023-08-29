package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookletFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 2v20H4v-4H2v-2h2v-3H2v-2h2V8H2V6h2V2h4Zm12.005 0C21.107 2 22 2.898 22 3.99v16.02c0 1.099-.893 1.99-1.995 1.99H10V2h10.005Z"/>`),
		g.Group(children),
	)
}