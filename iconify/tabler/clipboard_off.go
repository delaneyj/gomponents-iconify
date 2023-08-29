package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.575 5.597A2 2 0 0 0 5 7v12a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2m0-4V7a2 2 0 0 0-2-2h-2"/><path d="M9 5a2 2 0 0 1 2-2h2a2 2 0 1 1 0 4h-2M3 3l18 18"/></g>`),
		g.Group(children),
	)
}