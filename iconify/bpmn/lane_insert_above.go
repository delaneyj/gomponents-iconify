package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaneInsertAbove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M95 196v182h100v-82h32V196H95zm232 0v100h200V196H327zm300 0v100h200V196H627zm300 0v100h200V196H927zm300 0v100h200V196h-200zm300 0v100h200V196h-200zm298 0v100h26v82h100V196h-126zM95 478v188h100V478H95zm1756 0v188h100V478h-100zM96 762.643v1080h1856v-1080H96zm1756 97.699v891.543H192V860.344l1660-.002z"/>`),
		g.Group(children),
	)
}