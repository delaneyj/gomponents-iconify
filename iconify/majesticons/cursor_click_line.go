package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorClickLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M7.684 2.051a1 1 0 0 1 1.265.633l1 3a1 1 0 0 1-1.898.632l-1-3a1 1 0 0 1 .633-1.265zm8.023 2.242a1 1 0 0 1 0 1.414l-2 2a1 1 0 1 1-1.414-1.414l2-2a1 1 0 0 1 1.414 0zM2.051 7.683a1 1 0 0 1 1.265-.632l3 1a1 1 0 1 1-.632 1.898l-3-1a1 1 0 0 1-.633-1.265zm7.242 1.61a1 1 0 0 1 1.049-.233l11 4a1 1 0 0 1 .03 1.868l-3.593 1.437l3.928 3.928a1 1 0 0 1-1.414 1.414l-3.928-3.928l-1.436 3.592a1 1 0 0 1-1.869-.03l-4-11a1 1 0 0 1 .233-1.048zm2.38 2.38l2.371 6.523l1.027-2.567a1 1 0 0 1 .558-.557l2.567-1.027l-6.524-2.373zm-3.966.62a1 1 0 0 1 0 1.414l-2 2a1 1 0 1 1-1.414-1.414l2-2a1 1 0 0 1 1.414 0z"/></g>`),
		g.Group(children),
	)
}