package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrowningFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2"/><path stroke-width="3" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke-linecap="round" stroke-width="2" d="M8.535 16A3.998 3.998 0 0 1 12 14c1.48 0 2.773.804 3.465 2"/></g>`),
		g.Group(children),
	)
}