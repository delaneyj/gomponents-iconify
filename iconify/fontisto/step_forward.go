package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 22.835V1.146A1.141 1.141 0 0 1 1.696.149L1.69.146L20.4 10.639V.601a.6.6 0 0 1 .599-.6h1.2a.6.6 0 0 1 .6.6v22.8A.599.599 0 0 1 22.2 24h-1.201a.599.599 0 0 1-.599-.599V13.362L1.69 23.856a1.1 1.1 0 0 1-.548.145A1.155 1.155 0 0 1 0 22.846v-.011v.001z"/>`),
		g.Group(children),
	)
}