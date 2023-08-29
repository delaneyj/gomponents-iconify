package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoiceOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M31 24V11a7 7 0 1 0-14 0v13a7 7 0 1 0 14 0Z"/><path stroke-linecap="round" d="M9 23c0 8.284 6.716 15 15 15c1.753 0 3.436-.3 5-.853M39 23a14.95 14.95 0 0 1-1.248 6M24 38v6m18-2L6 6"/></g>`),
		g.Group(children),
	)
}