package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M15 26.314s4.596-.354 7.778-3.536C25.96 19.596 26.314 15 26.314 15l7.752 12.182a4.986 4.986 0 0 1-6.884 6.884L15 26.314Z"/><circle cx="15" cy="15" r="11"/><path stroke-linecap="round" stroke-linejoin="round" d="M5.657 25.456L25.456 5.657M34 33l8 8h-9"/></g>`),
		g.Group(children),
	)
}