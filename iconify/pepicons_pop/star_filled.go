package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="m10 15.97l-4.295 1.915a1 1 0 0 1-1.402-1.018l.494-4.677L1.65 8.697a1 1 0 0 1 .535-1.647l4.6-.976L9.134 2a1 1 0 0 1 1.732 0l2.35 4.074l4.6.976a1 1 0 0 1 .535 1.647l-3.148 3.494l.494 4.676a1 1 0 0 1-1.402 1.018L10 15.972Z"/>`),
		g.Group(children),
	)
}