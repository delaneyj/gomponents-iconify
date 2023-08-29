package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feBeer0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBeer1" fill="currentColor"><path id="feBeer2" d="M17 16v3c0 1.1-.9 2-2 2H5c-1.1 0-2-.9-2-2V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v2h2a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-2ZM5 5v3h10V5H5Zm0 5v9h10v-9H5Zm4 6a1 1 0 0 1-2 0v-3a1 1 0 0 1 2 0v3Zm4 0a1 1 0 0 1-2 0v-3a1 1 0 0 1 2 0v3Zm4-7h2v5h-2V9Z"/></g></g>`),
		g.Group(children),
	)
}