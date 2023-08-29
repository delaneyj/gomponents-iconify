package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnionSelection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M32 16h9a2 2 0 0 1 2 2v23a2 2 0 0 1-2 2H18a2 2 0 0 1-2-2v-9"/><path d="M32 16V7a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v23a2 2 0 0 0 2 2h9m25-16L17 40M32 7L7 32m25-16L16 32m27-8L24 43m0-38L5 24m38 10l-9 9M14 5l-9 9"/></g>`),
		g.Group(children),
	)
}