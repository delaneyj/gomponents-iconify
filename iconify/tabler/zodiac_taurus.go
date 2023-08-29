package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZodiacTaurus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 3a6 6 0 0 0 12 0"/><path d="M6 15a6 6 0 1 0 12 0a6 6 0 1 0-12 0"/></g>`),
		g.Group(children),
	)
}