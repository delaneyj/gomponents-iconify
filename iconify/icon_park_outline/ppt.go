package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ppt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 8h40"/><path d="M8 8h32v26H8V8Z" clip-rule="evenodd"/><path d="m22 16l5 5l-5 5m-6 16l8-8l8 8"/></g>`),
		g.Group(children),
	)
}