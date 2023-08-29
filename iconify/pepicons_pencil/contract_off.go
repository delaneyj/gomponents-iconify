package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContractOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M5.354 15.354a.5.5 0 0 1-.708-.708l4-4a.5.5 0 0 1 .708.708l-4 4Z"/><path d="M9.5 15a.5.5 0 0 1-1 0v-4a.5.5 0 0 1 1 0v4Z"/><path d="M5 11.5a.5.5 0 0 1 0-1h4a.5.5 0 0 1 0 1H5Zm6.354-2.146a.5.5 0 0 1-.708-.708l4-4a.5.5 0 0 1 .708.708l-4 4Z"/><path d="M11 9.5a.5.5 0 0 1 0-1h4a.5.5 0 0 1 0 1h-4Z"/><path d="M11.5 9a.5.5 0 0 1-1 0V5a.5.5 0 0 1 1 0v4ZM1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}