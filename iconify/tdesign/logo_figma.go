package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoFigma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5A4.5 4.5 0 0 1 8.5.5h7a4.5 4.5 0 0 1 2.829 8A4.5 4.5 0 0 1 13 15.742V19a4.5 4.5 0 1 1-7.329-3.5A4.491 4.491 0 0 1 4 12c0-1.414.652-2.675 1.671-3.5A4.491 4.491 0 0 1 4 5Zm4.5 2.5H11v-5H8.5a2.5 2.5 0 0 0 0 5Zm4.5-5v5h2.5a2.5 2.5 0 0 0 0-5H13Zm-2 7H8.5a2.5 2.5 0 0 0 0 5H11v-5Zm0 7H8.5A2.5 2.5 0 1 0 11 19v-2.5Zm2-4.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0Z"/>`),
		g.Group(children),
	)
}