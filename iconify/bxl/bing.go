package bxl

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.71 3l3.593 1.264v12.645l5.061-2.919l-2.48-1.165l-1.566-3.897l7.974 2.802v4.073l-8.984 5.183l-3.595-2L5.71 3z"/>`),
		g.Group(children),
	)
}