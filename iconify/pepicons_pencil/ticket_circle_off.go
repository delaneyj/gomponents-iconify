package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M4.5 9A1.5 1.5 0 0 1 6 7.5h14A1.5 1.5 0 0 1 21.5 9v2.64c0 .269-.161.51-.408.613a.809.809 0 0 0 0 1.493a.664.664 0 0 1 .408.613V17a1.5 1.5 0 0 1-1.5 1.5H6A1.5 1.5 0 0 1 4.5 17v-2.64c0-.269.161-.51.408-.614a.809.809 0 0 0 0-1.493a.664.664 0 0 1-.408-.612V9ZM6 8.5a.5.5 0 0 0-.5.5v2.431c1.208.683 1.208 2.455 0 3.138V17a.5.5 0 0 0 .5.5h14a.5.5 0 0 0 .5-.5v-2.431c-1.208-.683-1.208-2.455 0-3.138V9a.5.5 0 0 0-.5-.5H6Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15.5 16.25a.75.75 0 0 1 .75.75v.5a.75.75 0 0 1-1.5 0V17a.75.75 0 0 1 .75-.75Zm0-8.5a.75.75 0 0 1 .75.75V9a.75.75 0 0 1-1.5 0v-.5a.75.75 0 0 1 .75-.75Zm0 2.75a.75.75 0 0 1 .75.75v.5a.75.75 0 0 1-1.5 0v-.5a.75.75 0 0 1 .75-.75Zm0 3a.75.75 0 0 1 .75.75v.5a.75.75 0 0 1-1.5 0v-.5a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}