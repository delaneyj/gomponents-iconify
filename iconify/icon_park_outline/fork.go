package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M37 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm-26 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm13 32a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" d="M11 12v3c0 7 13 10 13 17v4v-4c0-7 13-10 13-17v-3"/></g>`),
		g.Group(children),
	)
}