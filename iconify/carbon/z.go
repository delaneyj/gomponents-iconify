package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Z(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M20 9h-8v2h6l-6 10v2h8v-2h-6l6-10V9z" fill="currentColor"/>`),
		g.Group(children),
	)
}