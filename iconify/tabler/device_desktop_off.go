package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceDesktopOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1m-4 0H4a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1m3 16h10m-8-4v4m6-4v4M3 3l18 18"/>`),
		g.Group(children),
	)
}