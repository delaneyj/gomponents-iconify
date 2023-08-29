package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Baby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 15h4s-1 1.5-2 1.5s-2-1.5-2-1.5Zm2-9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm6 6l-4-3m0 6l3 2.5l-2.5 2.5M6 12l4-3m0 6l-2.5 2.75L10 20m0-11h4v3h-4V9Z"/>`),
		g.Group(children),
	)
}