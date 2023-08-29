package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M17.5 1.5h-15a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h15a2 2 0 0 0 2-2v-10a2 2 0 0 0-2-2Zm-16 2a1 1 0 0 1 1-1h15a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-15a1 1 0 0 1-1-1v-10Z" clip-rule="evenodd"/><path d="M10 14a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5Z"/><path fill-rule="evenodd" d="M11.5 14.5h-3a1 1 0 0 0-1 1V18a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-2.5a1 1 0 0 0-1-1Zm-3 3.5v-2.5h3V18h-3Z" clip-rule="evenodd"/><path d="M5.5 19a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1h-9Z"/><path fill-rule="evenodd" d="M19 12H1v-1h18v1Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}