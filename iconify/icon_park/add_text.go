package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M4 8H32"/><path d="M28 21H44"/><path d="M18 42L18 8"/><path d="M36 42L36 21"/></g>`),
		g.Group(children),
	)
}