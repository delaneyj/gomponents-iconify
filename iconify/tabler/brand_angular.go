package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandAngular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m5.428 17.245l6.076 3.471a1 1 0 0 0 .992 0l6.076-3.471a1 1 0 0 0 .495-.734l1.323-9.704a1 1 0 0 0-.658-1.078l-7.4-2.612a1 1 0 0 0-.665 0L4.268 5.73a1 1 0 0 0-.658 1.078l1.323 9.704a1 1 0 0 0 .495.734z"/><path d="m9 15l3-8l3 8m-5-2h4"/></g>`),
		g.Group(children),
	)
}