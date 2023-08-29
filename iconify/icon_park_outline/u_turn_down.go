package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UTurnDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M35 34V15c0-6.075-4.925-11-11-11v0c-6.075 0-11 4.925-11 11v29"/><path stroke-linecap="round" stroke-linejoin="round" d="m18 39l-5 5l-5-5"/><circle cx="35" cy="39" r="5" transform="rotate(90 35 39)"/></g>`),
		g.Group(children),
	)
}