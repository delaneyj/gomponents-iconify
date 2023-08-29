package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceImacOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 3h13a1 1 0 0 1 1 1v12c0 .28-.115.532-.3.713M17 17H4a1 1 0 0 1-1-1V4c0-.276.112-.526.293-.707M3 13h10m4 0h4M8 21h8m-6-4l-.5 4m4.5-4l.5 4M3 3l18 18"/>`),
		g.Group(children),
	)
}