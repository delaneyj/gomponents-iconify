package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottleTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M15 21.56a9.82 9.82 0 0 1 4.767-8.42a.479.479 0 0 0 .233-.411V4h8v8.729c0 .168.088.324.233.41A9.82 9.82 0 0 1 33 21.56V42a2 2 0 0 1-2 2H17a2 2 0 0 1-2-2V21.56Z"/><path d="M20 10h8"/><path stroke-linejoin="round" d="M33 23h-9v15h9m0 2V21m-13-9V8m8 4V8"/></g>`),
		g.Group(children),
	)
}