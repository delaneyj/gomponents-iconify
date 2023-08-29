package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailPush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M36 15h8v26H4V15h8m12 4V5m6 6l-6-6l-6 6"/><path d="m4 15l20 15l20-15"/></g>`),
		g.Group(children),
	)
}