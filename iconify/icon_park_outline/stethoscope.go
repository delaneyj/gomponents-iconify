package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stethoscope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M21.947 4v6M10.053 7H4v14c0 5 4 11 12 11s12-6 12-11V7H10.053Zm0-3v6v-6ZM40 23a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M16 32c0 6.627 5.373 12 12 12s12-5.373 12-12v-9"/></g>`),
		g.Group(children),
	)
}