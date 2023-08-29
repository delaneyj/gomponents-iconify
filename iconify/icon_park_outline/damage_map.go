package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DamageMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M5 10v28a2 2 0 0 0 2 2h11l-3-11l7-2l-1-7l8-4l-2-3l3-5H7a2 2 0 0 0-2 2Zm38 28V10a2 2 0 0 0-2-2h-3l-4 6l3 5l-9 4l1 8l-7 2l2 7h17a2 2 0 0 0 2-2Z"/><path d="M14.5 18a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}