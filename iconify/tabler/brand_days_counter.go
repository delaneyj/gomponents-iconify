package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandDaysCounter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.779 10.007a9 9 0 1 0-10.77 10.772M13 21h8v-7"/><path d="M12 8v4l3 3"/></g>`),
		g.Group(children),
	)
}