package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M31 10L14 20m20-3L14 29m20-3L15 38M9 4h30v6H9zm0 34h30v6H9zm5-28v28"/><path d="M34 10v28"/></g>`),
		g.Group(children),
	)
}