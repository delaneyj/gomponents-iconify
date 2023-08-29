package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Railway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M12 6a2 2 0 0 1 2-2h20a2 2 0 0 1 2 2v24a2 2 0 0 1-2 2H14a2 2 0 0 1-2-2V6Z"/><circle cx="18" cy="26" r="2" fill="currentColor"/><circle cx="30" cy="26" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 20h24M21 38h6m-9 6h12"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="m17 32l-6 12m20-12l6 12"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 15v10M12 15v10"/></g>`),
		g.Group(children),
	)
}