package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCardLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m212.24 83.76l-56-56A6 6 0 0 0 152 26H56a14 14 0 0 0-14 14v176a14 14 0 0 0 14 14h144a14 14 0 0 0 14-14V88a6 6 0 0 0-1.76-4.24ZM202 216a2 2 0 0 1-2 2H56a2 2 0 0 1-2-2V40a2 2 0 0 1 2-2h93.52L202 90.49ZM74 120v72a6 6 0 0 0 6 6h96a6 6 0 0 0 6-6v-72a6 6 0 0 0-6-6H80a6 6 0 0 0-6 6Zm12 6h84v60h-20v-34a6 6 0 0 0-12 0v34h-20v-34a6 6 0 0 0-12 0v34H86Z"/>`),
		g.Group(children),
	)
}