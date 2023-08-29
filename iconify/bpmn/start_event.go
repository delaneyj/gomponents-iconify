package bpmn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StartEvent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="100.042" d="M1899 1023.999c0-483.252-391.75-874.995-875-874.995S149 540.747 149 1024c0 483.251 391.75 875.004 875 875.004s875-391.753 875-875.004z"/>`),
		g.Group(children),
	)
}