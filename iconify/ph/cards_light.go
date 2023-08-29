package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardsLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 74H40a14 14 0 0 0-14 14v112a14 14 0 0 0 14 14h144a14 14 0 0 0 14-14V88a14 14 0 0 0-14-14Zm2 126a2 2 0 0 1-2 2H40a2 2 0 0 1-2-2V88a2 2 0 0 1 2-2h144a2 2 0 0 1 2 2Zm44-144v120a6 6 0 0 1-12 0V56a2 2 0 0 0-2-2H64a6 6 0 0 1 0-12h152a14 14 0 0 1 14 14Z"/>`),
		g.Group(children),
	)
}