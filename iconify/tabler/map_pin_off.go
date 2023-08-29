package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPinOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9.442 9.432a3 3 0 0 0 4.113 4.134M15 11a3 3 0 0 0-3-3"/><path d="M17.152 17.162L13.414 20.9a2 2 0 0 1-2.827 0l-4.244-4.243a8 8 0 0 1-.476-10.794m2.18-1.82a8.003 8.003 0 0 1 10.91 10.912M3 3l18 18"/></g>`),
		g.Group(children),
	)
}