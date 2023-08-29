package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M16 12L4 24.432L16 36m16-24l12 12.432L32 36"/><path d="M24 17v14"/><path stroke-linejoin="round" d="m18 25l6 6l6-6"/></g>`),
		g.Group(children),
	)
}