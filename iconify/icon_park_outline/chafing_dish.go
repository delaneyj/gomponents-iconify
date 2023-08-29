package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChafingDish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M24 42c-9.941 0-18-8.059-18-18S14.059 6 24 6"/><path d="M24 42c9.941 0 18-8.059 18-18S33.941 6 24 6c0 0-4 2-4 9s8 11 8 18s-4 9-4 9Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M42 20h2v8h-2M6 20H4v8h2"/></g>`),
		g.Group(children),
	)
}