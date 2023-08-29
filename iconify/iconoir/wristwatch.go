package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wristwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16 16.472V20a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-3.528m0-8.944V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v3.528"/><path d="M18 12a6 6 0 1 0-12 0a6 6 0 0 0 12 0Z"/><path d="M14 12h-2v-2"/></g>`),
		g.Group(children),
	)
}