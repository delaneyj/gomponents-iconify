package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoutingBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="m20 19l.53.53a.75.75 0 0 0 0-1.06L20 19ZM8 4.25a.75.75 0 0 0 0 1.5v-1.5Zm10.53 12.22a.75.75 0 1 0-1.06 1.06l1.06-1.06Zm-1.06 4a.75.75 0 1 0 1.06 1.06l-1.06-1.06Zm-3.72-2.22a.75.75 0 0 0 0 1.5v-1.5ZM10 19.75a.75.75 0 0 0 0-1.5v1.5Zm1-14a.75.75 0 0 0 0-1.5v1.5Zm4-1.5a.75.75 0 0 0 0 1.5v-1.5Zm-7.5 8.5h9v-1.5h-9v1.5Zm13.03 5.72l-2-2l-1.06 1.06l2 2l1.06-1.06Zm-1.06 0l-2 2l1.06 1.06l2-2l-1.06-1.06Zm1.28-9.97a4.25 4.25 0 0 0-4.25-4.25v1.5a2.75 2.75 0 0 1 2.75 2.75h1.5Zm-4.25 4.25a4.25 4.25 0 0 0 4.25-4.25h-1.5a2.75 2.75 0 0 1-2.75 2.75v1.5ZM4.75 15.5a2.75 2.75 0 0 1 2.75-2.75v-1.5a4.25 4.25 0 0 0-4.25 4.25h1.5Zm2.75 2.75a2.75 2.75 0 0 1-2.75-2.75h-1.5a4.25 4.25 0 0 0 4.25 4.25v-1.5Zm12.5 0h-6.25v1.5H20v-1.5Zm-10 0H7.5v1.5H10v-1.5Zm1-14H8v1.5h3v-1.5Zm5.5 0H15v1.5h1.5v-1.5Z"/><circle cx="6" cy="5" r="2" stroke="currentColor" stroke-width="1.5"/></g>`),
		g.Group(children),
	)
}