package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PathExclude(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 5h10v4H9v6H5V5Zm4 10v4h10V9h-4v6H9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}