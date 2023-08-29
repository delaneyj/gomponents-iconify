package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationOffLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M11 22h2"/><circle cx="12" cy="3" r="1"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 19v-9c0-1.144.32-2.214.876-3.124M6 19h12M6 19H4m14 0v-1m0 1h1M9.999 4.342A6 6 0 0 1 18 10v2.343"/><path stroke-linecap="round" d="m4 4l16 16"/></g>`),
		g.Group(children),
	)
}