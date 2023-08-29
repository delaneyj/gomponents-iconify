package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MassageChair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="38" height="38" x="5" y="5" rx="3"/><path d="M18 24v-8.348C18 14.435 19.2 12 24 12s6 2.435 6 3.652V24m-14 0v6h16v-6"/><path d="M12 20v4h24v-4M18 36h12m-6-6v6"/></g>`),
		g.Group(children),
	)
}