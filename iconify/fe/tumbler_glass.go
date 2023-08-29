package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TumblerGlass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feTumblerGlass0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTumblerGlass1" fill="currentColor" fill-rule="nonzero"><path id="feTumblerGlass2" d="M7 3h10a2 2 0 0 1 2 2v10a6 6 0 0 1-6 6h-2a6 6 0 0 1-6-6V5a2 2 0 0 1 2-2Zm0 2v9h10V5H7Z"/></g></g>`),
		g.Group(children),
	)
}