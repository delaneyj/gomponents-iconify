package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func X(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M20 9h-2l-2 6l-2-6h-2l2.75 7L12 23h2l2-6l2 6h2l-2.75-7L20 9z" fill="currentColor"/>`),
		g.Group(children),
	)
}