package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandGoogleBigQuery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.73 19.875A2.225 2.225 0 0 1 15.782 21H8.499a2.222 2.222 0 0 1-1.947-1.158l-4.272-6.75a2.269 2.269 0 0 1 0-2.184l4.272-6.75A2.225 2.225 0 0 1 8.498 3h7.285c.809 0 1.554.443 1.947 1.158l3.98 6.75a2.33 2.33 0 0 1 0 2.25l-3.98 6.75v-.033z"/><path d="M8 11.5a3.5 3.5 0 1 0 7 0a3.5 3.5 0 1 0-7 0m6 2.5l2 2"/></g>`),
		g.Group(children),
	)
}