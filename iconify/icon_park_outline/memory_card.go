package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MemoryCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 40H8a2 2 0 0 1-2-2V19.106a2 2 0 0 1 .336-1.11l6.07-9.105A2 2 0 0 1 14.07 8H40a2 2 0 0 1 2 2v28a2 2 0 0 1-2 2ZM18 16V8m6 8V8m6 8V8m6 8V8"/><path d="M15 28h18v12H15V28Z"/></g>`),
		g.Group(children),
	)
}