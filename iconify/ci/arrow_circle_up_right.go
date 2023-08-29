package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 13V9m0 0h-4m4 0l-6 6m12-3a9 9 0 1 0-18 0a9 9 0 0 0 18 0Z"/>`),
		g.Group(children),
	)
}