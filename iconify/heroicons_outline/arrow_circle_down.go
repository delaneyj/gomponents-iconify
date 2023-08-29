package heroicons_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 13l-3 3m0 0l-3-3m3 3V8m0 13a9 9 0 1 1 0-18a9 9 0 0 1 0 18Z"/>`),
		g.Group(children),
	)
}