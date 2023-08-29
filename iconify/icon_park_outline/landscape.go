package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Landscape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M14 28c3 0 12.5 1.5 12.5 5s-5.106 2.739-7.5 3c-2.073 0-7 0-7 3c0 4 20.545 5 24.5 5"/><path d="M31 15c0-4.267-4.27-11-10.195-11c-5.925 0-10.472 8-12.26 15.467C6.757 26.933 7.013 36 7.013 36M26 24c0-2.5 1.4-7 6-7c4 0 6.2 4.042 7 7.5c.8 3.458 1 8.5 1 8.5"/></g>`),
		g.Group(children),
	)
}