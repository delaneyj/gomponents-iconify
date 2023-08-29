package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screenshot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M10 42a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path stroke-linecap="round" d="M40.062 8C24 28.433 15.805 38.68 14.242 40.243a6 6 0 0 1-8.485 0"/><path d="M38 42a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path stroke-linecap="round" d="M42.242 40.243a6 6 0 0 1-8.485 0C32.195 38.68 24 28.446 8 8"/></g>`),
		g.Group(children),
	)
}