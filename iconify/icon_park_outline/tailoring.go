package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tailoring(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M12 4v32h32"/><path stroke-linejoin="round" d="M20 12h16v16"/><path d="M12 12H4m32 32v-8"/></g>`),
		g.Group(children),
	)
}