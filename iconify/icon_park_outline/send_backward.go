package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M6 14h36V6H6zm0 14h36v-8H6z"/><path stroke-linecap="round" d="m30 36l-6 6l-6-6v0m6 6V28m0-14v6"/></g>`),
		g.Group(children),
	)
}