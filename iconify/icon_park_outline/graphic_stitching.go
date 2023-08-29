package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphicStitching(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-rule="evenodd"><path d="m24 24l10-10c0-5.523-4.477-10-10-10S14 8.477 14 14l10 10Z"/><path d="m14 34l10-10l-10-10C8.477 14 4 18.477 4 24s4.477 10 10 10Zm20 0c5.523 0 10-4.477 10-10s-4.477-10-10-10L24 24l10 10Z"/><path d="M24 44c5.523 0 10-4.477 10-10L24 24L14 34c0 5.523 4.477 10 10 10Z"/></g>`),
		g.Group(children),
	)
}