package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopRightBottomLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.17 8.66L21 11.5V3h-8.5l2.84 2.83l-9.51 9.51L3 12.5V21h8.5l-2.84-2.83l9.51-9.51Z"/>`),
		g.Group(children),
	)
}