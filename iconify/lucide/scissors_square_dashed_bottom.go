package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsSquareDashedBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2m-10 0H8m8 0h-2"/><circle cx="8" cy="8" r="2"/><path d="M9.414 9.414L12 12m2.8 2.8L18 18"/><circle cx="8" cy="16" r="2"/><path d="m18 6l-8.586 8.586"/></g>`),
		g.Group(children),
	)
}