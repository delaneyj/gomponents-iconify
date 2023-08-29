package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphicDesignTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M15.65 28.238A13.98 13.98 0 0 1 10 17c0-7.732 6.268-14 14-14s14 6.268 14 14c0 4.535-2.157 8.567-5.5 11.125"/><path d="m24 17l20 27H4l20-27Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}