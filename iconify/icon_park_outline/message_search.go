package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M25.5 37H21l-10 5v-5H4V7h40v11"/><circle cx="34" cy="28" r="6"/><path stroke-linecap="round" stroke-linejoin="round" d="m39 32l5 4M12 15h6m-6 6h12"/></g>`),
		g.Group(children),
	)
}