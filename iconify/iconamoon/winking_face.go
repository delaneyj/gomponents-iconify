package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WinkingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2"/><path stroke-width="3" d="M9 9.5h.01v.01H9z"/><path stroke-linecap="round" stroke-width="2" d="M15.465 14A3.998 3.998 0 0 1 12 16a3.998 3.998 0 0 1-3.465-2M16.5 9.5h-3"/></g>`),
		g.Group(children),
	)
}