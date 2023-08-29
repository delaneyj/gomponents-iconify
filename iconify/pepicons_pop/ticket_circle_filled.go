package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopTicketCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M4 9a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2.64c0 .47-.282.894-.716 1.075a.309.309 0 0 0 0 .57c.434.18.716.604.716 1.074V17a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-2.64c0-.47.282-.894.716-1.075a.309.309 0 0 0 0-.57A1.164 1.164 0 0 1 4 11.641V9Zm16 0H6v2.161c1.208.904 1.208 2.774 0 3.678V17h14v-2.161c-1.208-.905-1.208-2.774 0-3.678V9Z"/><path d="M15.5 16a1 1 0 0 1 1 1v.5a1 1 0 1 1-2 0V17a1 1 0 0 1 1-1Zm0-8.5a1 1 0 0 1 1 1V9a1 1 0 1 1-2 0v-.5a1 1 0 0 1 1-1Zm0 2.75a1 1 0 0 1 1 1v.5a1 1 0 1 1-2 0v-.5a1 1 0 0 1 1-1Zm0 3a1 1 0 0 1 1 1v.5a1 1 0 1 1-2 0v-.5a1 1 0 0 1 1-1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopTicketCircleFilled0)"/></g>`),
		g.Group(children),
	)
}