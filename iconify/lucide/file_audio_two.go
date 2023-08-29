package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileAudioTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v2"/><path d="M14 2v6h6M2 17v-3a4 4 0 0 1 8 0v3"/><circle cx="9" cy="17" r="1"/><circle cx="3" cy="17" r="1"/></g>`),
		g.Group(children),
	)
}