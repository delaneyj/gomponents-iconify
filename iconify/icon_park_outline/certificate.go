package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Certificate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M26 36H6a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v26a2 2 0 0 1-2 2h-8M12 14h24m-24 7h6m-6 7h4"/><path d="M30 33a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m30 40l4 2V31.472S32.86 33 30 33c-2.86 0-4-1.5-4-1.5V42l4-2Z"/></g>`),
		g.Group(children),
	)
}