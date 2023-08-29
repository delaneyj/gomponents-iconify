package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Garage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M6 5h36v38H6z"/><path stroke-linejoin="round" d="M12 12h24v6H12z"/><path d="M16 18v25m16-25v25M16 24h16m-16 6h16m-16 6h16"/></g>`),
		g.Group(children),
	)
}