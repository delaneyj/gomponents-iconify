package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailDelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 24V9H4v30h20"/><path d="m4 9l20 15L44 9M32 31h10l-2 10h-6l-2-10Zm4 0l2-4"/></g>`),
		g.Group(children),
	)
}