package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M17 4h14v30H17z"/><path stroke-linecap="round" d="M17 14h14M17 24h14M6.879 7.879l4.242 4.242m-4.242 21l4.242-4.243m30-20.999L36.88 12.12m4.241 21.001l-4.242-4.243M4 21h6m28 0h6"/><path d="M20 34h8v10h-8z"/></g>`),
		g.Group(children),
	)
}