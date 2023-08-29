package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Performance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M11 9v4.17a3.001 3.001 0 1 0 2 0V9h-2Zm0 7a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z" clip-rule="evenodd"/><path d="M12 5a7 7 0 0 1 7 7v1h-2v-1a5 5 0 0 0-10 0v1H5v-1a7 7 0 0 1 7-7Z"/><path fill-rule="evenodd" d="M12 23c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11Zm0-2a9 9 0 1 0 0-18a9 9 0 0 0 0 18Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}