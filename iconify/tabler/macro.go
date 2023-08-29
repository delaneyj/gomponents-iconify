package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Macro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 15a6 6 0 1 0 12 0"/><path d="M18 15a6 6 0 0 0-6 6m0 0a6 6 0 0 0-6-6m6 6V11m0 0a5 5 0 0 1-5-5V3l3 2l2-2l2 2l3-2v3a5 5 0 0 1-5 5z"/></g>`),
		g.Group(children),
	)
}