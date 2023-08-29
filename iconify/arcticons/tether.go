package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tether(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.358 40.964h-6.911a1 1 0 0 1-1-1V23.931a.067.067 0 0 0-.067-.067H5.862a.384.384 0 0 1-.257-.668L23.256 7.32a1.112 1.112 0 0 1 1.487 0l17.652 15.875a.384.384 0 0 1-.257.668h-3.585v16.1a1 1 0 0 1-1 1h-6.91"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.642 40.964h-9.401v-6.178a2.83 2.83 0 1 1 5.662 0m-.001 0v2.077m1.316-7.781a5.199 5.199 0 0 0-4.125-1.745a5.199 5.199 0 0 0-4.126 1.745m11.397-3.447a9.976 9.976 0 0 0-7.282-3.399a9.675 9.675 0 0 0-7.27 3.163"/>`),
		g.Group(children),
	)
}