package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrowningFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path stroke-width="1.5" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5z"/><path stroke-linecap="round" d="M8.535 16A3.998 3.998 0 0 1 12 14c1.48 0 2.773.804 3.465 2"/></g>`),
		g.Group(children),
	)
}