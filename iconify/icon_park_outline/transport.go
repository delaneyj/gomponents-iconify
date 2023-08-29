package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><rect width="28" height="18" x="16" y="12" stroke-linejoin="round" rx="3"/><path stroke-linecap="round" d="M24 18v6m12-6v6"/><path stroke-linecap="round" stroke-linejoin="round" d="M36 12V6H24v6m20 24H12a2 2 0 0 1-2-2V11a2 2 0 0 0-2-2H4m15 33a3 3 0 0 1-3-3v-3h6v3a3 3 0 0 1-3 3Zm18 0a3 3 0 0 1-3-3v-3h6v3a3 3 0 0 1-3 3Z"/></g>`),
		g.Group(children),
	)
}