package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refraction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m24 9l16.454 28.5H7.545L24 9Z"/><path d="m24 9l16.454 28.5H7.545L24 9ZM4 22l15.5-5m8.5-1l16-3m-14 6.5L44 21m-11.3 3L44 29"/></g>`),
		g.Group(children),
	)
}