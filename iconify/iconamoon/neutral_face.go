package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NeutralFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke-linejoin="round" stroke-width="3" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke-linecap="round" stroke-width="2" d="M9 15h6"/></g>`),
		g.Group(children),
	)
}