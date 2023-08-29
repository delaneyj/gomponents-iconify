package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M200 64v16h-24a16 16 0 0 0-16 16v80H40a16 16 0 0 1-16-16V64a16 16 0 0 1 16-16h144a16 16 0 0 1 16 16Z" opacity=".2"/><path d="M224 72h-16v-8a24 24 0 0 0-24-24H40a24 24 0 0 0-24 24v96a24 24 0 0 0 24 24h112v8a24 24 0 0 0 24 24h48a24 24 0 0 0 24-24V96a24 24 0 0 0-24-24ZM40 168a8 8 0 0 1-8-8V64a8 8 0 0 1 8-8h144a8 8 0 0 1 8 8v8h-16a24 24 0 0 0-24 24v72Zm192 24a8 8 0 0 1-8 8h-48a8 8 0 0 1-8-8V96a8 8 0 0 1 8-8h48a8 8 0 0 1 8 8Zm-96 16a8 8 0 0 1-8 8H88a8 8 0 0 1 0-16h40a8 8 0 0 1 8 8Zm80-96a8 8 0 0 1-8 8h-16a8 8 0 0 1 0-16h16a8 8 0 0 1 8 8Z"/></g>`),
		g.Group(children),
	)
}