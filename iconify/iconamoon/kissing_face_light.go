package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KissingFaceLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="1.5"/><path stroke-linecap="round" stroke-width="1.5" d="m13 16l-2-1l2-1"/><path stroke-width="2.25" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5z"/></g>`),
		g.Group(children),
	)
}