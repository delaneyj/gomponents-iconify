package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func H(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M18 9v6h-4V9h-2v14h2v-6h4v6h2V9h-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}