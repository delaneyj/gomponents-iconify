package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feBus0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBus1" fill="currentColor" fill-rule="nonzero"><path id="feBus2" d="M16 19H8a2 2 0 1 1-4 0V5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v14a2 2 0 1 1-4 0Zm1-4a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM7 15a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM6 5v6h12V5H6Z"/></g></g>`),
		g.Group(children),
	)
}