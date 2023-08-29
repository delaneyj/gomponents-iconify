package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableStackBelowTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.5 9.5v5h5v-5h-5ZM8 9.5v5H3.75a.75.75 0 0 1-.75-.75V9.5h5ZM9.5 8h5V3h-5v5ZM16 9.5v5h4.25a.75.75 0 0 0 .75-.75V9.5h-5ZM21 8h-5V3h1.75A3.25 3.25 0 0 1 21 6.25V8ZM8 8H3V6.25A3.25 3.25 0 0 1 6.25 3H8v5ZM3.75 19.5a.75.75 0 0 0 0 1.5h16.5a.75.75 0 0 0 0-1.5H3.75Z"/>`),
		g.Group(children),
	)
}