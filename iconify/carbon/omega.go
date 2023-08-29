package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Omega(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22.737 25A14.31 14.31 0 0 0 27 15c0-6.58-4.42-11-11-11S5 8.42 5 15a14.31 14.31 0 0 0 4.263 10H4v2h8v-2.238A12.568 12.568 0 0 1 7 15c0-5.467 3.533-9 9-9s9 3.533 9 9a12.568 12.568 0 0 1-5 9.762V27h8v-2Z"/>`),
		g.Group(children),
	)
}