package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandDoctrine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 14a7 7 0 1 0 14 0a7 7 0 1 0-14 0m4 0h6"/><path d="m12 11l3 3l-3 3M10 3l6.9 6"/></g>`),
		g.Group(children),
	)
}