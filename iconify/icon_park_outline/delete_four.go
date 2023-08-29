package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M8 11h32M18 5h12"/><path d="M12 17h24v23a3 3 0 0 1-3 3H15a3 3 0 0 1-3-3V17Z"/></g>`),
		g.Group(children),
	)
}