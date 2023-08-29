package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumbersTwoOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 8a5 5 0 0 1 5-5h1v.1c2.282.463 4 2.481 4 4.9v.065c0 1.525-.687 2.97-1.871 3.931L9 16.976V18h8v2H7v-3.024a2 2 0 0 1 .739-1.552l6.129-4.98A3.065 3.065 0 0 0 15 8.065V8a3 3 0 1 0-6 0v1H7V8Z"/>`),
		g.Group(children),
	)
}