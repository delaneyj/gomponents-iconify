package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 3.672h2V16.5l3.243-3.243l1.414 1.415L12 20.328l-5.657-5.656l1.414-1.415L11 16.5V3.672Z"/>`),
		g.Group(children),
	)
}