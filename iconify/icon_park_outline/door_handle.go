package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorHandle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="26" height="40" x="6" y="4" rx="2"/><path d="M14 34h10m18-14v-6H23a5 5 0 1 0 0 6h19Z"/></g>`),
		g.Group(children),
	)
}