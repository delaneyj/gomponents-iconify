package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glasses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="14" r="4"/><circle cx="18" cy="14" r="4"/><path d="m10 14l.211-.106a4 4 0 0 1 3.578 0L14 14m5-8l2.838 6.623a2 2 0 0 1 .162.788V14M5 6l-2.838 6.623A2 2 0 0 0 2 13.41V14"/></g>`),
		g.Group(children),
	)
}