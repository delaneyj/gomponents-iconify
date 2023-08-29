package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Database(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5 9V7h2v2H5Zm4 0h10V7H9v2Zm-4 6v2h2v-2H5Zm14 2H9v-2h10v2Z"/><path fill-rule="evenodd" d="M1 6a3 3 0 0 1 3-3h16a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V6Zm3-1h16a1 1 0 0 1 1 1v5H3V6a1 1 0 0 1 1-1Zm-1 8v5a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-5H3Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}