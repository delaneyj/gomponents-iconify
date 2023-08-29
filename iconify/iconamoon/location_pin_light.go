package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationPinLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="2.25" d="M12 11h.01v.01H12z"/><path stroke-width="1.5" d="m12 22l5.5-5.5a7.778 7.778 0 1 0-11 0L12 22Z"/></g>`),
		g.Group(children),
	)
}