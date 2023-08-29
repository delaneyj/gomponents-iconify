package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingZoomAreaZoomMagnifierSquareArea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.5.5h1a1 1 0 0 1 1 1v1m-13 0v-1a1 1 0 0 1 1-1h1m3 0h3m5 5v2.75M.5 5.5v3m0 3v1a1 1 0 0 0 1 1h1m3 0h2.75"/><circle cx="8" cy="8" r="3.5"/><path d="M10.47 10.47L13 13"/></g>`),
		g.Group(children),
	)
}