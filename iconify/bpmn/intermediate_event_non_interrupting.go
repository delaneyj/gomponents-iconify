package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IntermediateEventNonInterrupting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="100" transform="translate(0 995.638)"><circle cx="1024" cy="28.357" r="875" stroke-dasharray="418.31 361.233"/><circle cx="1024" cy="28.357" r="685" stroke-dasharray="348.31 261.233" stroke-dashoffset="500"/></g>`),
		g.Group(children),
	)
}