package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PciCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M0 1.5A.5.5 0 0 1 .5 1h1a.5.5 0 0 1 .5.5V4h13.5a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-.5.5H2v2.5a.5.5 0 0 1-1 0V2H.5a.5.5 0 0 1-.5-.5Z"/><path d="M3 12.5h3.5v1a.5.5 0 0 1-.5.5H3.5a.5.5 0 0 1-.5-.5v-1Zm4 0h4v1a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-1Z"/></g>`),
		g.Group(children),
	)
}