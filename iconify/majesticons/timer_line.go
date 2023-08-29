package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 7l-1.343 1.343m0 0A8 8 0 1 0 6.343 19.657A8 8 0 0 0 17.657 8.343zM12 10v4M9 3h6"/>`),
		g.Group(children),
	)
}