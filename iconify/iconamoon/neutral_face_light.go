package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NeutralFaceLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/><path stroke-linejoin="round" stroke-width="2.25" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5z"/><path stroke-linecap="round" stroke-width="1.5" d="M9 15h6"/></g>`),
		g.Group(children),
	)
}