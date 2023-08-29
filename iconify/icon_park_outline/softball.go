package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Softball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M24 44c11.05 0 20-8.95 20-20S35.05 4 24 4S4 12.95 4 24s8.95 20 20 20Z"/><path stroke-linecap="round" d="M10 38c3.7-3.63 6-8.41 6-14c0-5.52-2.38-10.38-6-14m28 28c-3.7-3.63-6-8.41-6-14c0-5.52 2.38-10.38 6-14M28 24h8m-24 0h8m10 10l6-3m0-14l-6-3M18 34l-6-3m0-14l6-3"/></g>`),
		g.Group(children),
	)
}