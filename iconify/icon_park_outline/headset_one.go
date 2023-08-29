package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadsetOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="M36 32a8 8 0 1 0 0-16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M36 32a8 8 0 1 0 0-16"/><path d="M12 16a8 8 0 1 0 0 16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M12 16a8 8 0 1 0 0 16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 32V16c0-6.627 5.373-12 12-12s12 5.373 12 12v16c0 6.627-5.373 12-12 12"/></g>`),
		g.Group(children),
	)
}