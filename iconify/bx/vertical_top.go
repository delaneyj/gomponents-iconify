package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 7l-4 4h3v8h2v-8h3l-4-4zM3 3h3v2H3zm5 0h3v2H8zm5 0h3v2h-3zm5 0h3v2h-3z"/>`),
		g.Group(children),
	)
}