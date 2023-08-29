package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Needle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M39 23L25 9m19 8L31 4m3.999 14.999l5-5m-11-1l5-5M11 37l-5 5"/><path d="M28 12L12 28l-1 9l9-1l16-16l-8-8Z"/></g>`),
		g.Group(children),
	)
}