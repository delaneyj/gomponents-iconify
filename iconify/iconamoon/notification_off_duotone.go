package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M6 10v9h12v-1L6.876 6.876A5.972 5.972 0 0 0 6 10Z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 22h2"/><circle cx="12" cy="3" r="1" stroke="currentColor" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.999 4.342A6 6 0 0 1 18 10v2.343M6 19v-9c0-1.144.32-2.214.876-3.124M6 19h12M6 19H4m14 0v-1m0 1h1"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m4 4l16 16"/></g>`),
		g.Group(children),
	)
}