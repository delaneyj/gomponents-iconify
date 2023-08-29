package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Text(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M2.5 4.5h37v9h-2.85l-2.15-7h-10v27l6 1.8v2.2h-19v-2.2l6-1.8v-27H7.55l-2.2 7H2.5z"/>`),
		g.Group(children),
	)
}