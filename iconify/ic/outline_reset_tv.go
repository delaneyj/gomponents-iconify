package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineResetTv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 8V5c0-1.1-.9-2-2-2H4c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h4v2h8v-2h4c1.1 0 1.99-.9 1.99-2v-5H22c0-1.1-.9-2-2-2h-7.17l1.83-1.83l-1.41-1.41C9.69 10.31 10.88 9.12 9 11l4.24 4.24l1.41-1.41L12.83 12H20v5H4V5h16v3h2z"/>`),
		g.Group(children),
	)
}