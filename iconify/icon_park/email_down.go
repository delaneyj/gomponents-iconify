package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M36 15H44V28V41H4V28V15H12"/><path d="M24 19V5"/><path d="M30 13L24 19L18 13"/><path d="M4 15L24 30L44 15"/></g>`),
		g.Group(children),
	)
}