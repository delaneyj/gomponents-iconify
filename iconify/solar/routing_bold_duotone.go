package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoutingBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M17.47 16.47a.75.75 0 0 1 1.06 0l2 2a.75.75 0 0 1 0 1.06l-2 2a.75.75 0 1 1-1.06-1.06L18.94 19l-1.47-1.47a.75.75 0 0 1 0-1.06Z" clip-rule="evenodd"/><path d="M8 5a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/><path d="M7.855 5.75a1.994 1.994 0 0 0 0-1.5H16.5a4.25 4.25 0 0 1 0 8.5h-9a2.75 2.75 0 1 0 0 5.5h10.69l.75.75l-.75.75H7.5a4.25 4.25 0 0 1 0-8.5h9a2.75 2.75 0 1 0 0-5.5H7.855Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}