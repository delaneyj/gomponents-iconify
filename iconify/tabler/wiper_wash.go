package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WiperWash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 20a1 1 0 1 0 2 0a1 1 0 1 0-2 0m-8-9l5.5 5.5a5 5 0 0 1 7 0L21 11a12 12 0 0 0-18 0m9 9V6M4 6a4 4 0 0 1 .4-1.8M7 2.1a4 4 0 0 1 2 0"/><path d="M12 6a4 4 0 0 0-.4-1.8"/><path d="M12 6a4 4 0 0 1 .4-1.8M15 2.1a4 4 0 0 1 2 0M20 6a4 4 0 0 0-.4-1.8"/></g>`),
		g.Group(children),
	)
}