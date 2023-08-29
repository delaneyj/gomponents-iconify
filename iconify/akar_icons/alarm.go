package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="13" r="9"/><path d="M15.5 9.5L12 13m7 6l1 3M5 19l-1 3M2 5l3-3m14 0l3 3M12 4V2"/></g>`),
		g.Group(children),
	)
}