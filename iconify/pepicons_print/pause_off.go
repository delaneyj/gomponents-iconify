package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path d="M7 7a2 2 0 1 1 4 0v8a2 2 0 1 1-4 0V7Z"/><path fill-rule="evenodd" d="M11.5 7v8a2.5 2.5 0 0 1-5 0V7a2.5 2.5 0 0 1 5 0ZM9 5a2 2 0 0 0-2 2v8a2 2 0 1 0 4 0V7a2 2 0 0 0-2-2Z" clip-rule="evenodd"/><path d="M13 7a2 2 0 1 1 4 0v8a2 2 0 1 1-4 0V7Z"/><path fill-rule="evenodd" d="M17.5 7v8a2.5 2.5 0 0 1-5 0V7a2.5 2.5 0 0 1 5 0ZM15 5a2 2 0 0 0-2 2v8a2 2 0 1 0 4 0V7a2 2 0 0 0-2-2Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M8 14V6a1 1 0 0 0-2 0v8a1 1 0 1 0 2 0ZM7 4a2 2 0 0 0-2 2v8a2 2 0 1 0 4 0V6a2 2 0 0 0-2-2Zm7 10V6a1 1 0 1 0-2 0v8a1 1 0 1 0 2 0ZM13 4a2 2 0 0 0-2 2v8a2 2 0 1 0 4 0V6a2 2 0 0 0-2-2Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}