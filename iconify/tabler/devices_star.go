package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 13V9a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v2.5"/><path d="M18 8V5a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h5.5m8.3 2.817l-2.172 1.138a.392.392 0 0 1-.568-.41l.415-2.411l-1.757-1.707a.389.389 0 0 1 .217-.665l2.428-.352l1.086-2.193a.392.392 0 0 1 .702 0l1.086 2.193l2.428.352a.39.39 0 0 1 .217.665l-1.757 1.707l.414 2.41a.39.39 0 0 1-.567.411L17.8 20.817zM16 9h2"/></g>`),
		g.Group(children),
	)
}