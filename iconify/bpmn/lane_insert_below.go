package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaneInsertBelow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M95 1842.643v-182h100v82h32v100H95zm232 0v-100h200v100H327zm300 0v-100h200v100H627zm300 0v-100h200v100H927zm300 0v-100h200v100h-200zm300 0v-100h200v100h-200zm298 0v-100h26v-82h100v182h-126zm-1730-282v-188h100v188H95zm1756 0v-188h100v188h-100zM96 1276V196h1856v1080H96zm1756-97.7V286.759H192v891.54l1660 .003z"/>`),
		g.Group(children),
	)
}