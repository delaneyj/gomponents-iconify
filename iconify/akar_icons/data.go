package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Data(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="6" rx="8" ry="3"/><path d="M6.037 12C4.77 12.53 4 13.232 4 14c0 1.657 3.582 3 8 3s8-1.343 8-3c0-.768-.77-1.47-2.037-2"/><path d="M4 6v4c0 1.657 3.582 3 8 3s8-1.343 8-3V6M4 14v4c0 1.657 3.582 3 8 3s8-1.343 8-3v-4"/></g>`),
		g.Group(children),
	)
}