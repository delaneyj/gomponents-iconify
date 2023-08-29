package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenJacket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m6 10l12-6h12l12 6l-2 25h-6v9H14v-9H8L6 10Zm8 25V20m20 15V20"/><path d="M24 10a6 6 0 0 0 6-6H18a6 6 0 0 0 6 6Z"/></g>`),
		g.Group(children),
	)
}