package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 6C14.059 6 6 14.034 6 23.944V26c0-2.478 2.015-5 4.5-5s4.5 2.522 4.5 5c0-2.478 2.015-5 4.5-5s4.5 2.522 4.5 5c0-2.478 2.015-5 4.5-5s4.5 2.522 4.5 5c0-2.478 2.015-5 4.5-5s4.5 2.522 4.5 5v-2.056C42 14.034 33.941 6 24 6Z"/><path d="M15 26s-1.5-5.5 1-11c2.501-5.5 8-9 8-9m9 20s.5-4.5-2-11s-7-9-7-9m0 19v14.5a4.5 4.5 0 0 1-4.5 4.5v0a4.5 4.5 0 0 1-4.5-4.5v-1.318M24 25V6m0 0V4"/><path d="M28.5 21c2.485 0 4.5 2.522 4.5 5c0-2.478 2.015-5 4.5-5m-9 0c-2.485 0-4.5 2.522-4.5 5c0-2.478-2.015-5-4.5-5m-9 0c2.485 0 4.5 2.522 4.5 5c0-2.478 2.015-5 4.5-5M15.733 8C18.21 6.722 21.02 6 24 6c2.98 0 5.79.722 8.266 2"/></g>`),
		g.Group(children),
	)
}