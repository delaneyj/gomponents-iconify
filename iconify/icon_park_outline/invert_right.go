package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InvertRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M38 15c0 7.299-4.103 13.583-10 16.408A16.147 16.147 0 0 1 21 33c-9.389 0-17-8.059-17-18"/><path d="m30 20l8-5l6 8"/></g>`),
		g.Group(children),
	)
}