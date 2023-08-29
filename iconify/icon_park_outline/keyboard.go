package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path fill="currentColor" fill-rule="evenodd" d="M15 19a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0 8a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm9-8a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0 8a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm9-8a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0 8a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 33h14"/></g>`),
		g.Group(children),
	)
}