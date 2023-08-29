package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M4.5 18L24 8l20 10m-20 0V4M10 30V16"/><circle cx="10" cy="36" r="6"/><path stroke-linecap="round" stroke-linejoin="round" d="M38 30V16m-6 20l6-6l6 6l-6 6l-6-6Zm-3-13v-5H19v5"/></g>`),
		g.Group(children),
	)
}