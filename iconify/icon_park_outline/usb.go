package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M12 22a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm24 6a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" d="m19 9l5-5l5 5m-4 30L12 28.263V22m24 6v4.79L24 41m0-37v39m-3 1h6"/></g>`),
		g.Group(children),
	)
}