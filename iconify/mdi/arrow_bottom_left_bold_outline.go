package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBottomLeftBoldOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.36 19.78H4.22V5.64l4.25 4.24l5.65-5.66l5.66 5.66l-5.66 5.66l4.24 4.24M6.34 17.66h7.07l-2.12-2.12l5.66-5.66l-2.83-2.83l-5.66 5.66l-2.12-2.12v7.07Z"/>`),
		g.Group(children),
	)
}