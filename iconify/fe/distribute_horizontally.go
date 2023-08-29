package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeHorizontally(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feDistributeHorizontally0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDistributeHorizontally1" fill="currentColor"><path id="feDistributeHorizontally2" d="M7 21a1 1 0 0 1-2 0V3a1 1 0 1 1 2 0v18Zm12 0a1 1 0 0 1-2 0V3a1 1 0 0 1 2 0v18Zm-5-4a2 2 0 1 1-4 0V7a2 2 0 1 1 4 0v10Z"/></g></g>`),
		g.Group(children),
	)
}