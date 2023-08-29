package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsExpandDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M11 5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V5ZM9 5H5v4h4V5Z" clip-rule="evenodd"/><path d="M19 13h2v8h-8v-2h4.586l-5.364-5.364a1 1 0 0 1 1.414-1.414L19 17.586V13Z"/></g>`),
		g.Group(children),
	)
}