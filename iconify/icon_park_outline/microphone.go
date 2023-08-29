package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><rect width="16" height="28" x="16" y="4" stroke-linecap="round" stroke-linejoin="round" rx="8"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 21v3c0 7.732 6.268 14 14 14v0c7.732 0 14-6.268 14-14v-3M24 5v6m-8 5h5m6 0h5m-16 6h5m6 0h5"/><path d="M24 38v6"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 44h16"/></g>`),
		g.Group(children),
	)
}