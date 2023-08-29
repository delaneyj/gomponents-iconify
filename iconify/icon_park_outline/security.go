package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Security(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 26v-4c0-9.941-8.059-18-18-18v0C14.059 4 6 12.059 6 22v4"/><path d="M32 27V12a8 8 0 0 0-8-8v0a8 8 0 0 0-8 8v15"/><path d="M24 24v14a6 6 0 0 0 6 6v0a6 6 0 0 0 6-6v-2.889M12 24h30"/></g>`),
		g.Group(children),
	)
}