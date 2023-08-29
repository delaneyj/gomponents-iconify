package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Receive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="70"><path d="M224 500h1552v1000H224z"/><path stroke-miterlimit="1.4" d="m240 500l760 500l760-500z"/></g>`),
		g.Group(children),
	)
}