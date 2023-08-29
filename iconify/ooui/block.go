package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Block(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 1a9 9 0 1 0 9 9a9 9 0 0 0-9-9zm5 10H5V9h10z"/>`),
		g.Group(children),
	)
}