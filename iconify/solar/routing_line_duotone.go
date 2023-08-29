package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoutingLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M8 4.25a.75.75 0 0 0 0 1.5v-1.5Zm12 14H7.5v1.5H20v-1.5Zm-12.5-5.5h9v-1.5h-9v1.5Zm9-8.5H8v1.5h8.5v-1.5Zm4.25 4.25a4.25 4.25 0 0 0-4.25-4.25v1.5a2.75 2.75 0 0 1 2.75 2.75h1.5Zm-4.25 4.25a4.25 4.25 0 0 0 4.25-4.25h-1.5a2.75 2.75 0 0 1-2.75 2.75v1.5ZM4.75 15.5a2.75 2.75 0 0 1 2.75-2.75v-1.5a4.25 4.25 0 0 0-4.25 4.25h1.5Zm2.75 2.75a2.75 2.75 0 0 1-2.75-2.75h-1.5a4.25 4.25 0 0 0 4.25 4.25v-1.5Z" opacity=".5"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m18 17l2 2l-2 2"/><circle cx="6" cy="5" r="2" stroke="currentColor" stroke-width="1.5"/></g>`),
		g.Group(children),
	)
}