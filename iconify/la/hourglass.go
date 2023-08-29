package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7 4v2h2v4a7.001 7.001 0 0 0 3.406 6A7.001 7.001 0 0 0 9 22v4H7v2h18v-2h-2v-4a7.001 7.001 0 0 0-3.406-6A7.001 7.001 0 0 0 23 10V6h2V4zm4 2h10v4c0 2.773-2.227 5-5 5s-5-2.227-5-5zm5 11c2.773 0 5 2.227 5 5v4H11v-4c0-2.773 2.227-5 5-5z"/>`),
		g.Group(children),
	)
}