package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Towel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M36 18H4v8h32v-8Z"/><path d="M36 12v20a4 4 0 0 1-4 4H12m0 0H8a4 4 0 0 1-4-4V8a4 4 0 0 1 4-4h32a4 4 0 0 1 4 4v32c0 2.21-1.79 4-4 4H16c-2.21 0-4-1.79-4-4v-4Z"/></g>`),
		g.Group(children),
	)
}