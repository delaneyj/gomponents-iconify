package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewsDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 4h14v16H5a2 2 0 0 1-2-2V4Z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8h-4"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4h14v14a2 2 0 0 0 2 2v0M13 8H7m6 4H9"/></g>`),
		g.Group(children),
	)
}