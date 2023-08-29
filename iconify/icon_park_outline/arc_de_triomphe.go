package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArcDeTriomphe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M8 16v28h10V29a6 6 0 0 1 12 0v15h10V16H8Zm-3-6h38v3a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3v-3Z"/><path stroke-linecap="round" d="M5 4h38M8 4v5m32-5v5M8 28h10m12 0h10"/></g>`),
		g.Group(children),
	)
}