package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrangeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M34.15 34.85c8.202-8.202 8.202-21.498 0-29.7L4.453 34.85c8.201 8.2 21.498 8.2 29.699 0ZM19.301 20l14.85 14.85M19.302 20v21.213m0-21.213l20.506.707"/><path d="M39.755 14.997c1.664 6.88-.189 14.437-5.56 19.808c-5.372 5.372-12.93 7.225-19.809 5.56m-.034-15.415l8.485-8.486"/></g>`),
		g.Group(children),
	)
}