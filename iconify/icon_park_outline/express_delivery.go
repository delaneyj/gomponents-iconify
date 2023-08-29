package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpressDelivery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M8 31v11a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V31"/><path d="M38 14H10a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2Z"/><path stroke-linecap="round" d="M16 4v4m8-4v4m8-4v4M16 34h16"/></g>`),
		g.Group(children),
	)
}