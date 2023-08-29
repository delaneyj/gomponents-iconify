package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M17 7h-7a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-7"/><path d="M17 4h14v6H17z"/></g>`),
		g.Group(children),
	)
}