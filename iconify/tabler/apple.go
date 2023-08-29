package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 14a7 7 0 1 0 14 0a7 7 0 1 0-14 0"/><path d="M12 11V5a2 2 0 0 1 2-2h2v1a2 2 0 0 1-2 2h-2"/><path d="M10 10.5c1.333.667 2.667.667 4 0"/></g>`),
		g.Group(children),
	)
}