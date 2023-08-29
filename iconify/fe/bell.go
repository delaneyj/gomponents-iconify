package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feBell0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBell1" fill="currentColor"><path id="feBell2" d="M15 18v1a3 3 0 0 1-6 0v-1H5c-.55 0-1-.45-1-1s.45-1 1-1h1v-6a6.002 6.002 0 0 1 5-5.917V3a1 1 0 0 1 2 0v1.083c2.838.476 5 2.944 5 5.917v6h1c.55 0 1 .45 1 1s-.45 1-1 1h-4Zm-7-2h8v-6a4 4 0 1 0-8 0v6Zm4 4a1 1 0 0 0 1-1v-1h-2v1a1 1 0 0 0 1 1Z"/></g></g>`),
		g.Group(children),
	)
}