package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M33 7.26261C30.3212 5.81915 27.2563 5 24 5C13.5066 5 5 13.5066 5 24C5 34.4934 13.5066 43 24 43C26.858 43 29.5685 42.369 32 41.2387"/><path stroke-linejoin="round" d="M31 30L43 30"/><path stroke-linejoin="round" d="M15 22L22 29L41 11"/><path stroke-linejoin="round" d="M37 24V36"/></g>`),
		g.Group(children),
	)
}