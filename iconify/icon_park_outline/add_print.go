package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M12 19H6V6h36v13h-6"/><path stroke-linejoin="round" d="M12 12h24v32l-6-4.444L24 44l-6-4.444L12 44V12Z"/><path d="M20 26h8m-4-4v8"/></g>`),
		g.Group(children),
	)
}