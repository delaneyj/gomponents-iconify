package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 8l-4 4l4 4v-3h8v-2h-8V8zM3 18h2v3H3zm0-5h2v3H3zm0-5h2v3H3zm0-5h2v3H3z"/>`),
		g.Group(children),
	)
}