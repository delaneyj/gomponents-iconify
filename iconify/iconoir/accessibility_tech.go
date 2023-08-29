package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessibilityTech(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 19V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m12.5 12.16l4-.16l-.5 4.5M11.833 12L13.5 9.538L10.833 8L9.5 9.846"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.5 7.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M10.5 18a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/></g>`),
		g.Group(children),
	)
}