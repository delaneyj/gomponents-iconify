package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationPinDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M17.5 16.5L12 22l-5.5-5.5a7.778 7.778 0 1 1 11 0ZM12 12a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 11h.01v.01H12z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m12 22l5.5-5.5a7.778 7.778 0 1 0-11 0L12 22Z"/></g>`),
		g.Group(children),
	)
}