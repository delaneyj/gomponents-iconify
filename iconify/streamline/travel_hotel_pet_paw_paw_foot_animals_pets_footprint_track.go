package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelHotelPetPawPawFootAnimalsPetsFootprintTrack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="1.5" cy="8" rx="1" ry="1.5"/><ellipse cx="4.5" cy="3.5" rx="1" ry="1.5"/><ellipse cx="9.5" cy="3.5" rx="1" ry="1.5"/><ellipse cx="12.5" cy="8" rx="1" ry="1.5"/><path d="M10 10c0 1.38-1.62 2-3 2s-3-.62-3-2s1-3.5 3-3.5s3 2.12 3 3.5Z"/></g>`),
		g.Group(children),
	)
}