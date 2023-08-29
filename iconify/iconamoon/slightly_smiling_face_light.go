package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlightlySmilingFaceLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="1.5"/><path stroke-width="2.25" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5z"/><path stroke-linecap="round" stroke-width="1.5" d="M15.465 14A3.999 3.999 0 0 1 12 16a3.998 3.998 0 0 1-3.465-2"/></g>`),
		g.Group(children),
	)
}