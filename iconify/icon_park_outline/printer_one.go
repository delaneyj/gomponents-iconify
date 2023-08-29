package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M32 33h10a2 2 0 0 0 2-2V15a2 2 0 0 0-2-2H32M16 33H6a2 2 0 0 1-2-2V15a2 2 0 0 1 2-2h10"/><path d="M12 13v6h24v-6m-20 0V5h16v8M16 29v14h16V29m-10 6h4m-13-8h22"/></g>`),
		g.Group(children),
	)
}