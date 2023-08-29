package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Union(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feUnion0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUnion1" fill="currentColor"><path id="feUnion2" d="M15 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h4V5a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-4Zm-2-2h6V5h-8v6H5v8h8v-6Z"/></g></g>`),
		g.Group(children),
	)
}