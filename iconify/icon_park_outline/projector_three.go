package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="M6 38v-6h36v6h-6v-6H12v6H6Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 38v-6h-6v6h6ZM6 38v-6h6v6H6Zm18-22H4v16h40V16h-6m-28 8h8"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M31 23a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/><path fill="currentColor" d="M31 19a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`),
		g.Group(children),
	)
}