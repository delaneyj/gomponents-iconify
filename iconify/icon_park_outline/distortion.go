package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Distortion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 19v21h13M4 19V8h13M4 19s9-6 20 0s20 0 20 0m0 0V8H17m27 11v21H17M4 29s9-6 20 0s20 0 20 0M17 8s6 8 0 16s0 16 0 16M31 8s6 8 0 16s0 16 0 16"/>`),
		g.Group(children),
	)
}