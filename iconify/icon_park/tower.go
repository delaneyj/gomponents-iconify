package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="14" height="13" x="17" y="31" fill="#2F88FF"/><rect width="10" height="14" x="19" y="17" fill="#2F88FF"/><rect width="6" height="13" x="21" y="4" fill="#2F88FF"/><path stroke-linecap="round" d="M4 44H44"/></g>`),
		g.Group(children),
	)
}