package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDBridge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 4h3"/><path fill="currentColor" d="M10 21a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4-16a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M10 20s6.5-2.5 2-8s2-8 2-8M3 20h3"/></g>`),
		g.Group(children),
	)
}