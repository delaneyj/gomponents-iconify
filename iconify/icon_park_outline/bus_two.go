package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 6a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v32a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6Z" clip-rule="evenodd"/><path d="M16 40H8v4h8v-4Zm24 0h-8v4h8v-4ZM21 16h6M10 34h2m7 0h10M4 25h40M4 10h40m-8 24h2M4 6v8m40-8v8M4 21v8m40-8v8"/></g>`),
		g.Group(children),
	)
}