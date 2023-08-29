package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M1 6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2.64c0 .47-.282.894-.716 1.075a.309.309 0 0 0 0 .57c.434.18.716.604.716 1.074V14a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2v-2.64c0-.47.282-.894.716-1.075a.309.309 0 0 0 0-.57A1.164 1.164 0 0 1 1 8.641V6Zm16 0H3v2.161c1.208.904 1.208 2.774 0 3.678V14h14v-2.161c-1.208-.905-1.208-2.774 0-3.678V6Z"/><path d="M12.5 13a1 1 0 0 1 1 1v.5a1 1 0 1 1-2 0V14a1 1 0 0 1 1-1Zm0-8.5a1 1 0 0 1 1 1V6a1 1 0 1 1-2 0v-.5a1 1 0 0 1 1-1Zm0 2.75a1 1 0 0 1 1 1v.5a1 1 0 1 1-2 0v-.5a1 1 0 0 1 1-1Zm0 3a1 1 0 0 1 1 1v.5a1 1 0 1 1-2 0v-.5a1 1 0 0 1 1-1Z"/></g>`),
		g.Group(children),
	)
}