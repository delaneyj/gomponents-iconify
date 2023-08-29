package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Workbench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M12 33H4V7h40v26H12Z"/><path stroke-linecap="round" d="M16 22v4m8 7v6m0-21v8m8-12v12M12 41h24"/></g>`),
		g.Group(children),
	)
}