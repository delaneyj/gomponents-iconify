package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M15.244 21.366a2.164 2.164 0 0 1-3.061 0l-8.549-8.549A2.164 2.164 0 0 1 3 11.287V5.163C3 3.97 3.97 3 5.164 3h6.123c.573 0 1.124.228 1.53.634l8.549 8.549a2.164 2.164 0 0 1 0 3.061l-6.122 6.122Z"/><path d="M6.5 6.5L7 7"/></g>`),
		g.Group(children),
	)
}