package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapDistance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M39.3 6H8.7A2.7 2.7 0 0 0 6 8.7v30.6A2.7 2.7 0 0 0 8.7 42h30.6a2.7 2.7 0 0 0 2.7-2.7V8.7A2.7 2.7 0 0 0 39.3 6Z"/><path stroke-linecap="round" d="m36 27l-7 3m-8 3l-7 3"/><path d="M16 29c2-2.895 3-4.895 3-6a3 3 0 1 0-6 0c0 1.105 1 3.105 3 6Zm16-7c2-2.895 3-4.895 3-6a3 3 0 1 0-6 0c0 1.105 1 3.105 3 6Z"/></g>`),
		g.Group(children),
	)
}