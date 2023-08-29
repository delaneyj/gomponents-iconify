package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 7v5l2 2m3 8l5-3l-5-3z"/><path d="M13.017 20.943a9 9 0 1 1 7.831-7.292"/></g>`),
		g.Group(children),
	)
}