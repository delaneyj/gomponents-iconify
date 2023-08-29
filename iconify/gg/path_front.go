package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PathFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 5h10v4h4v10H9v-4H5V5Zm6 6v6h6v-6h-6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}