package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelHotelHangerHangerLockerCheckCoatRoomCloak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 3a2 2 0 1 0-2 2v1m0 0l-6 4.6a1.32 1.32 0 0 0-.5 1.06A1.34 1.34 0 0 0 1.84 13h10.32a1.34 1.34 0 0 0 1.34-1.34a1.32 1.32 0 0 0-.5-1.06Z"/>`),
		g.Group(children),
	)
}