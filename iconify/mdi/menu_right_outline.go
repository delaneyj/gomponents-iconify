package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuRightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 6h1.5l6 6l-6 6H9V6m4.67 6L11 9.33v5.34L13.67 12Z"/>`),
		g.Group(children),
	)
}