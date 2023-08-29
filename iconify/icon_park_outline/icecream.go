package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icecream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M33 18H15l6 22s1 3 3 3s3-3 3-3l6-22Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M36 18H12c0-8 5-14 12-14s12 6 12 14Z"/></g>`),
		g.Group(children),
	)
}