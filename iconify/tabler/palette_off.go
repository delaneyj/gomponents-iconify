package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaletteOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 15h-1a2 2 0 0 0-1 3.75A1.3 1.3 0 0 1 12 21A9 9 0 0 1 5.628 5.644M8 4c1.236-.623 2.569-1 4-1c4.97 0 9 3.582 9 8c0 1.06-.474 2.078-1.318 2.828a4.516 4.516 0 0 1-1.127.73"/><path d="M7.5 10.5a1 1 0 1 0 2 0a1 1 0 1 0-2 0m4-3a1 1 0 1 0 2 0a1 1 0 1 0-2 0m4 3a1 1 0 1 0 2 0a1 1 0 1 0-2 0M3 3l18 18"/></g>`),
		g.Group(children),
	)
}