package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListSuccess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M20 10h24M20 24h24M20 38h24"/><circle cx="8" cy="24" r="4"/><circle cx="8" cy="38" r="4"/><path d="m4 10l3 3l6-6"/></g>`),
		g.Group(children),
	)
}