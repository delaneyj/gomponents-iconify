package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M6.5 6a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 4a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 4a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/><path fill-rule="evenodd" d="M7.5 6.5A.5.5 0 0 1 8 6h8a.5.5 0 0 1 0 1H8a.5.5 0 0 1-.5-.5Zm0 4A.5.5 0 0 1 8 10h8a.5.5 0 0 1 0 1H8a.5.5 0 0 1-.5-.5Zm0 4A.5.5 0 0 1 8 14h8a.5.5 0 0 1 0 1H8a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}