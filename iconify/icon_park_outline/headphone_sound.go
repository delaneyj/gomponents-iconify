package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphoneSound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path d="M4 28a2 2 0 0 1 2-2h4v12H6a2 2 0 0 1-2-2v-8Zm34-2h4a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-4V26Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 36V24c0-7.732 6.268-14 14-14s14 6.268 14 14v12M10 26H6a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h4V26Zm28 0h4a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-4V26Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 32h4l2-6l4 12l2-6h4"/></g>`),
		g.Group(children),
	)
}