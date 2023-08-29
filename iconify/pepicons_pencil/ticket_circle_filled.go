package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilTicketCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M4.5 9A1.5 1.5 0 0 1 6 7.5h14A1.5 1.5 0 0 1 21.5 9v2.64c0 .269-.161.51-.408.613a.809.809 0 0 0 0 1.493a.664.664 0 0 1 .408.613V17a1.5 1.5 0 0 1-1.5 1.5H6A1.5 1.5 0 0 1 4.5 17v-2.64c0-.269.161-.51.408-.614a.809.809 0 0 0 0-1.493a.664.664 0 0 1-.408-.612V9ZM6 8.5a.5.5 0 0 0-.5.5v2.431c1.208.683 1.208 2.455 0 3.138V17a.5.5 0 0 0 .5.5h14a.5.5 0 0 0 .5-.5v-2.431c-1.208-.683-1.208-2.455 0-3.138V9a.5.5 0 0 0-.5-.5H6Z"/><path d="M15.5 16.25a.75.75 0 0 1 .75.75v.5a.75.75 0 0 1-1.5 0V17a.75.75 0 0 1 .75-.75Zm0-8.5a.75.75 0 0 1 .75.75V9a.75.75 0 0 1-1.5 0v-.5a.75.75 0 0 1 .75-.75Zm0 2.75a.75.75 0 0 1 .75.75v.5a.75.75 0 0 1-1.5 0v-.5a.75.75 0 0 1 .75-.75Zm0 3a.75.75 0 0 1 .75.75v.5a.75.75 0 0 1-1.5 0v-.5a.75.75 0 0 1 .75-.75Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilTicketCircleFilled0)"/></g>`),
		g.Group(children),
	)
}