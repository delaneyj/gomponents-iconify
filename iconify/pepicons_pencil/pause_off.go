package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M8 14V6a1 1 0 0 0-2 0v8a1 1 0 1 0 2 0ZM7 4a2 2 0 0 0-2 2v8a2 2 0 1 0 4 0V6a2 2 0 0 0-2-2Zm7 10V6a1 1 0 1 0-2 0v8a1 1 0 1 0 2 0ZM13 4a2 2 0 0 0-2 2v8a2 2 0 1 0 4 0V6a2 2 0 0 0-2-2Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}