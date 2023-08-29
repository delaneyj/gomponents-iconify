package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BringForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M6 34h36v8H6zm0-14h36v8H6z"/><path stroke-linecap="round" d="m30 12l-6-6l-6 6v0m6 16v6m0-28v14"/></g>`),
		g.Group(children),
	)
}