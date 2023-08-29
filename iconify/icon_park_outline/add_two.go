package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="M33 7.263A18.916 18.916 0 0 0 24 5C13.507 5 5 13.507 5 24s8.507 19 19 19a18.93 18.93 0 0 0 8-1.761"/><path stroke-linejoin="round" d="M31 30h12m-28-8l7 7l19-18m-4 13v12"/></g>`),
		g.Group(children),
	)
}