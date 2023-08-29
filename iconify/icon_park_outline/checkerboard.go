package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checkerboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M17 4H7a3 3 0 0 0-3 3v34a3 3 0 0 0 3 3h34a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3H17Zm6 13h21M4 17h9m22 14h9M6 31h19m-8-10v23M31 4v23m0 8v9M17 4v9"/><path d="M35 31a4 4 0 1 1-8 0a4 4 0 0 1 8 0ZM21 17a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/></g>`),
		g.Group(children),
	)
}