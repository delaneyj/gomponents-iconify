package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cooking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" d="M6 42h36M6 36h36"/><path stroke-linecap="round" stroke-linejoin="round" d="M9 25c0-8.284 6.716-15 15-15c8.284 0 15 6.716 15 15v11H9V25Z"/><path stroke-linecap="round" d="M17 25v4"/><path d="M28 10V8a4 4 0 0 0-8 0v2"/></g>`),
		g.Group(children),
	)
}