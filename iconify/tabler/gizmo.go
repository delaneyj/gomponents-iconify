package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gizmo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m20 19l-8-5.5L4 19m8-15v9.5M11 4a1 1 0 1 0 2 0a1 1 0 1 0-2 0"/><path d="M3 19a1 1 0 1 0 2 0a1 1 0 1 0-2 0m16 0a1 1 0 1 0 2 0a1 1 0 1 0-2 0"/></g>`),
		g.Group(children),
	)
}