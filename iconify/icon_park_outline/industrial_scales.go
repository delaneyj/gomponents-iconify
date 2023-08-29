package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndustrialScales(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M10 32h28l4 8H6l4-8Z"/><path stroke-linecap="round" d="M16 40v4m8-32v20"/><path d="M17 4h14v8H17z"/><path stroke-linecap="round" d="M32 40v4"/></g>`),
		g.Group(children),
	)
}