package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameController(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 15l-2.968 2.968A2.362 2.362 0 0 1 2 16.298V15l1.357-6.784A4 4 0 0 1 7.279 5h9.442a4 4 0 0 1 3.922 3.216L22 15v1.297a2.362 2.362 0 0 1-4.032 1.67L15 15H9Z"/><path d="m9 5l1 2h4l1-2"/></g>`),
		g.Group(children),
	)
}