package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowAutofitRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12V6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v8m0 4h17m-3-3l3 3l-3 3"/>`),
		g.Group(children),
	)
}