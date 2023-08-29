package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.5 6A1.5 1.5 0 0 1 3 4.5h14A1.5 1.5 0 0 1 18.5 6v2.64c0 .269-.161.51-.408.613a.809.809 0 0 0 0 1.493a.664.664 0 0 1 .408.613V14a1.5 1.5 0 0 1-1.5 1.5H3A1.5 1.5 0 0 1 1.5 14v-2.64c0-.269.161-.51.408-.614a.809.809 0 0 0 0-1.493a.664.664 0 0 1-.408-.612V6ZM3 5.5a.5.5 0 0 0-.5.5v2.431c1.208.683 1.208 2.455 0 3.138V14a.5.5 0 0 0 .5.5h14a.5.5 0 0 0 .5-.5v-2.431c-1.208-.683-1.208-2.455 0-3.138V6a.5.5 0 0 0-.5-.5H3Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.5 13.25a.75.75 0 0 1 .75.75v.5a.75.75 0 0 1-1.5 0V14a.75.75 0 0 1 .75-.75Zm0-8.5a.75.75 0 0 1 .75.75V6a.75.75 0 0 1-1.5 0v-.5a.75.75 0 0 1 .75-.75Zm0 2.75a.75.75 0 0 1 .75.75v.5a.75.75 0 0 1-1.5 0v-.5a.75.75 0 0 1 .75-.75Zm0 3a.75.75 0 0 1 .75.75v.5a.75.75 0 0 1-1.5 0v-.5a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}