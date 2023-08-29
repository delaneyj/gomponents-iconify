package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMicrophoneOne0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" d="M15 26.314s4.596-.354 7.778-3.536C25.96 19.596 26.314 15 26.314 15l7.752 12.182a4.986 4.986 0 0 1-6.884 6.884L15 26.314Z"/><circle cx="15" cy="15" r="11"/><path stroke-linecap="round" stroke-linejoin="round" d="M5.657 25.456L25.456 5.657M34 33l8 8h-9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMicrophoneOne0)"/>`),
		g.Group(children),
	)
}