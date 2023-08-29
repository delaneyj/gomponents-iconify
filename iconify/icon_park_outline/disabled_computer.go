package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisabledComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M23 5.998H9a3 3 0 0 0-3 3v22a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3v-7M24 34v8"/><circle cx="36" cy="12" r="6"/><path stroke-linecap="round" stroke-linejoin="round" d="m32 8l8 8M14 41.998h20"/></g>`),
		g.Group(children),
	)
}