package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandWindy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 4c0 5.5-.33 16 4 16s7.546-11.27 8-13"/><path d="M3 4c.253 5.44 1.449 16 5.894 16C13.338 20 17.314 9.964 18 6"/></g>`),
		g.Group(children),
	)
}