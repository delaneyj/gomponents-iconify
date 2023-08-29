package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BridgeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M8 13s6 10 16 10s16-10 16-10"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 10v28m32-28v28"/><path stroke-linecap="round" d="M4 30.5s12.188-.597 20-.5c7.82.098 20 1 20 1M16 21v9m8-7v7m8-9v9M8 13l-4 5m40 0l-4-5"/></g>`),
		g.Group(children),
	)
}