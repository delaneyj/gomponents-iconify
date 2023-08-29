package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m7 17l-5-5l5-5m5 10l-5-5l5-5"/><path d="M22 18v-2a4 4 0 0 0-4-4H7"/></g>`),
		g.Group(children),
	)
}