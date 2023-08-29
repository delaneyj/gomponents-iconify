package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceTabletFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 24H64a24 24 0 0 0-24 24v160a24 24 0 0 0 24 24h128a24 24 0 0 0 24-24V48a24 24 0 0 0-24-24ZM64 40h128a8 8 0 0 1 8 8v8H56v-8a8 8 0 0 1 8-8Zm128 176H64a8 8 0 0 1-8-8v-8h144v8a8 8 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}