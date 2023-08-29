package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardTab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 18h2V6h-2m-8.41 1.41L15.17 11H1v2h14.17l-3.58 3.58L13 18l6-6l-6-6l-1.41 1.41Z"/>`),
		g.Group(children),
	)
}