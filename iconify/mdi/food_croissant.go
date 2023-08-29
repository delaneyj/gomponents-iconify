package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodCroissant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 19l-3-2l3-2v4m-7-4l4-6l3 4l-4 3l-3-1M5 17l-3 2v-4l3 2m4-2l-3 1l-4-3l3-4l4 6m5-9l4 2l-5 7h-2L6 8l4-2h4Z"/>`),
		g.Group(children),
	)
}