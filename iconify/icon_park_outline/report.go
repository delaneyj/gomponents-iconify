package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Report(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M36 35H12V21c0-6.627 5.373-12 12-12s12 5.373 12 12v14Z"/><path stroke-linecap="round" d="M8 42h32M4 13l3 1m6-10l1 3m-4 3L7 7"/></g>`),
		g.Group(children),
	)
}