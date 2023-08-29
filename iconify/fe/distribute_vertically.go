package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeVertically(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feDistributeVertically0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDistributeVertically1" fill="currentColor"><path id="feDistributeVertically2" d="M7 10h10a2 2 0 1 1 0 4H7a2 2 0 1 1 0-4Zm-4 7h18a1 1 0 0 1 0 2H3a1 1 0 0 1 0-2ZM3 5h18a1 1 0 0 1 0 2H3a1 1 0 1 1 0-2Z"/></g></g>`),
		g.Group(children),
	)
}