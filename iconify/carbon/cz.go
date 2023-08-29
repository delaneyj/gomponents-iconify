package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M14 23H9a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h5v2H9v6h5z" fill="currentColor"/><path d="M24 9h-8v2h6l-6 10v2h8v-2h-6l6-10V9z" fill="currentColor"/>`),
		g.Group(children),
	)
}