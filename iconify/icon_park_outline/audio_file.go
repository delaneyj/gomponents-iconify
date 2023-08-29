package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AudioFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 44V4h23l9 10.5V44H8Z"/><path d="m32 14l-6 2.969V31.5"/><circle cx="20.5" cy="31.5" r="5.5"/></g>`),
		g.Group(children),
	)
}