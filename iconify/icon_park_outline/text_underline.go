package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextUnderline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M8 44h32"/><path d="M37 6.097V22c0 7.18-5.82 13-13 13s-13-5.82-13-13V6.097"/></g>`),
		g.Group(children),
	)
}