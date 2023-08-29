package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContractCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M8.354 18.354a.5.5 0 0 1-.708-.708l4-4a.5.5 0 0 1 .708.708l-4 4Z"/><path d="M12.5 18a.5.5 0 0 1-1 0v-4a.5.5 0 0 1 1 0v4Z"/><path d="M8 14.5a.5.5 0 0 1 0-1h4a.5.5 0 0 1 0 1H8Zm6.354-2.146a.5.5 0 0 1-.708-.708l4-4a.5.5 0 0 1 .708.708l-4 4Z"/><path d="M14 12.5a.5.5 0 0 1 0-1h4a.5.5 0 0 1 0 1h-4Z"/><path d="M14.5 12a.5.5 0 0 1-1 0V8a.5.5 0 0 1 1 0v4ZM4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}