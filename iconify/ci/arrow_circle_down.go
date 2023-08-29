package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 13l3 3m0 0l3-3m-3 3V8m9 4a9 9 0 1 0-18 0a9 9 0 0 0 18 0Z"/>`),
		g.Group(children),
	)
}