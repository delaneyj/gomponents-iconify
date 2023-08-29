package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PathOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 5h10v4h4v10H9v-4H5V5Zm2 2h6v2H9v4H7V7Zm4 10h6v-6h-2v4h-4v2Zm2-6h-2v2h2v-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}