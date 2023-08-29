package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandLaravel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 17l8 5l7-4v-8l-4-2.5L18 5l4 2.5v4L11 18l-4-2.5V8L3 5.5zm8 1v4m-4-6.5l7-4m0-4v4m0 0l4 2.5"/><path d="M11 13V5.5L7 3L3 5.5M7 8l4-2.5m7 4.5l4-2.5"/></g>`),
		g.Group(children),
	)
}