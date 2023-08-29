package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CandleF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 9h6v6H7V9zm-5 8h16a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-1a2 2 0 0 1 2-2zm8-8.6a3 3 0 0 1-3-3C7 4.295 8 2.495 10 0c2 2.495 3 4.295 3 5.4a3 3 0 0 1-3 3z"/>`),
		g.Group(children),
	)
}