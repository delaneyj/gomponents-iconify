package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Court(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 36h20V12H4v24h20Zm0 0v-8m0-16v8"/><circle cx="24" cy="24" r="4"/><path d="M11 24a4 4 0 0 1-4 4H4v-8h3a4 4 0 0 1 4 4Zm26 0a4 4 0 0 0 4 4h3v-8h-3a4 4 0 0 0-4 4Z"/></g>`),
		g.Group(children),
	)
}