package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareToggle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v20m2-2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h8m6 2a2 2 0 0 0-2-2m0 16a2 2 0 0 0 2-2m0-8v4"/>`),
		g.Group(children),
	)
}