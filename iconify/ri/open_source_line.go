package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenSourceLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 2c5.523 0 10 4.477 10 10c0 4.4-2.841 8.136-6.789 9.473l-.226.074l-2.904-7.55A2.016 2.016 0 0 0 14.001 12a2 2 0 1 0-2.083 1.998l-2.903 7.549l-.225-.074C4.842 20.136 2 16.4 2 12C2 6.477 6.477 2 12 2Zm0 2a8 8 0 0 0-4.099 14.872l1.48-3.849a4 4 0 1 1 5.239 0c.565 1.474 1.059 2.757 1.479 3.85A8 8 0 0 0 12 4Z"/>`),
		g.Group(children),
	)
}