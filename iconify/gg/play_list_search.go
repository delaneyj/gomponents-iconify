package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayListSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M15.879 4.879h-12v2h12v-2Zm0 4h-12v2h12v-2Zm-12 4h8v2h-8v-2Z"/><path fill-rule="evenodd" d="M13.757 12.757a3 3 0 0 0 3.415 4.83l1.535 1.534l1.414-1.414l-1.535-1.535a3.001 3.001 0 0 0-4.829-3.415Zm1.415 2.829a1 1 0 1 0 1.414-1.415a1 1 0 0 0-1.414 1.415Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}