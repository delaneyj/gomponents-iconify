package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompareVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 15v7h2v-7h3l-4-4l-4 4h3m-1-6h-3V2H8v7H5l4 4l4-4Z"/>`),
		g.Group(children),
	)
}