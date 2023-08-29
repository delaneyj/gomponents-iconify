package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FishHook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 9v6a5 5 0 0 1-10 0v-4l3 3"/><path d="M14 7a2 2 0 1 0 4 0a2 2 0 1 0-4 0m2-2V3"/></g>`),
		g.Group(children),
	)
}