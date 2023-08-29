package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowAutofitUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8m4 0V3m-3 3l3-3l3 3"/>`),
		g.Group(children),
	)
}