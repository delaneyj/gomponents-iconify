package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Juice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M15 24h18l-1.8 20H16.8L15 24Z"/><rect width="26" height="6" x="11" y="18" rx="3"/><path d="M24 8c-5.523 0-10 4.477-10 10h20c0-5.523-4.477-10-10-10Z"/><path stroke-linecap="round" d="m28 4l-2 4"/></g>`),
		g.Group(children),
	)
}