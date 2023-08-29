package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M25.5 19H33a2 2 0 0 1 2 2v23H13V21a2 2 0 0 1 2-2h2.5"/><path stroke-linejoin="round" d="M17 8h-2.5a5.5 5.5 0 1 0 0 11H19v10.5a2.5 2.5 0 0 0 5 0V19h9.5a5.5 5.5 0 1 0 0-11H29s-1-4-6-4s-6 4-6 4Z"/><path d="M35 21h5a2 2 0 0 1 2 2v5a4 4 0 0 1-4 4h-3"/></g>`),
		g.Group(children),
	)
}