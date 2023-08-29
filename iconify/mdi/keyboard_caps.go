package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardCaps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18h12v-2H6m6-7.59L16.59 13L18 11.58l-6-6l-6 6L7.41 13L12 8.41Z"/>`),
		g.Group(children),
	)
}