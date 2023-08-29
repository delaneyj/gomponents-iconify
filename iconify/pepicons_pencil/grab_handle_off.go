package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrabHandleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M2.5 7a.5.5 0 0 1 0-1h15a.5.5 0 0 1 0 1h-15Zm0 3.25a.5.5 0 0 1 0-1h15a.5.5 0 0 1 0 1h-15Zm0 3.25a.5.5 0 0 1 0-1h15a.5.5 0 0 1 0 1h-15Z"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}